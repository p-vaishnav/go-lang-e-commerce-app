package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ServiceUnavailable(ctx *gin.Context, err string) {
	var res models.BaseResponse
	errorCode := http.StatusServiceUnavailable

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Warn("SERVICE UNAVAILABLE",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
