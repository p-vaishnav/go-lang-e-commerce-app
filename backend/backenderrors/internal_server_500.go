package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServer(ctx *gin.Context, err string) {
	var res models.BaseResponse
	errorCode := http.StatusInternalServerError

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = "internal server error"
	res.Status = "error"

	// log error
	// logs.Error("SERVER",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
