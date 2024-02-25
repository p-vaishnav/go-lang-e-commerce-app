package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Unauthorized(ctx *gin.Context, err string, debugMessage string) {
	var res models.BaseResponse
	errorCode := http.StatusUnauthorized

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Info("UNAUTHORIZED",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
