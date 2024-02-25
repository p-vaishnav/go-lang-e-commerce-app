package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Forbidden(ctx *gin.Context, err string) {
	var res models.BaseResponse
	errorCode := http.StatusForbidden

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Info("FORBIDDEN",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
