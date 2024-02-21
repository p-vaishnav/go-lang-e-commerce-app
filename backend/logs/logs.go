package logs

import (
	"backend-commerce/constants"
	"backend-commerce/entities"
	"backend-commerce/services/authsvc"
	"backend-commerce/utils"
	"bytes"
	"io"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type responseBodyWriter struct {
	gin.ResponseWriter // what do these both serve
	body               *bytes.Buffer
}

func (r *responseBodyWriter) Write(b []byte) (int, error) {
	r.body.Write(b)
	return r.ResponseWriter.Write(b)
}

func findParams(ctx *gin.Context) []string {
	var params []string

	for _, param := range ctx.Params {
		params = append(params, param.Key+":"+param.Value)
	}

	return params
}

// TODO: go through gin.Context's fields

func RequestResponseLogs(ctx *gin.Context, gormDB *gorm.DB) {
	// ignore health check logs
	if ctx.Request.URL.Path == "/" {
		return
	}

	var reqBodyBytes []byte
	var userPID string
	var err error

	// should I write a test case for it
	if ctx.Request != nil {
		reqBodyBytes, err = io.ReadAll(ctx.Request.Body)
		if err != nil {
			// TODO: write a log over here
			return //
		}
		ctx.Request.Body = io.NopCloser(bytes.NewBuffer(reqBodyBytes)) // it helps to re-read the request body again and again
	}

	w := &responseBodyWriter{body: &bytes.Buffer{}, ResponseWriter: ctx.Writer} // writer interface implementation starts here
	ctx.Writer = w                                                              // didn't get this line and the above one as well

	ctx.Next() // to go and complete the rest of the work - what if inbetween panic happens ??

	params := strings.Join(findParams(ctx), ",")
	contentType := strings.ToLower(ctx.ContentType())
	traceID := ctx.GetHeader(constants.Headers.TraceID)

	signedToken := ctx.GetHeader("Authorization")

	if signedToken != "" {
		// its only for the protected routes only

		authSvc := authsvc.Handler()
		userPID, err = authSvc.VerifyToken(ctx, signedToken)
		if err != nil {
			// TODO: write a log over here
		}
	}

	remoteIP := ctx.RemoteIP()
	requestResponseLogs := &entities.RequestResponseLogs{
		PID:                utils.UUIDWithPrefix(constants.Prefix.Logs),
		UserPID:            userPID,
		Method:             ctx.Request.Method,
		ResponseCode:       ctx.Writer.Status(),
		RequestBodyLength:  ctx.Request.ContentLength,
		ResponseBodyLength: int64(len(w.body.String())),
		EndPoint:           ctx.Request.URL.Path,
		HostURL:            ctx.Request.Host,
		ClientIP:           ctx.ClientIP(),
		RemoteIP:           remoteIP,
		Params:             params,
		QueryParams:        ctx.Request.URL.RawQuery,
		ContentType:        contentType,
		TraceID:            traceID,
	}

	requestResponseLogs.RequestBody = entities.ReqResJson{reqBodyBytes}
	requestResponseLogs.ResponseBody = entities.ReqResJson{w.body.Bytes()}

	// TODO: check for the form updated

	txnDB := gormDB.Create(requestResponseLogs)
	if txnDB.Error != nil {
		// log some thing
		return
	}

}
