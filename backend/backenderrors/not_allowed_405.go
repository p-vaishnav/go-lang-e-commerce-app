package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func NotAllowed(ctx *gin.Context, err string) {
	var res models.BaseResponse
	errorCode := http.StatusMethodNotAllowed

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Info("NOT ALLOWED",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
