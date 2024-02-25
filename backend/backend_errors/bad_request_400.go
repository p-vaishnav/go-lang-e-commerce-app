package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func BadRequest(ctx *gin.Context, err string, debugMessage string) {
	var res models.BaseResponse
	errorCode := http.StatusBadRequest

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Info("BAD REQUEST",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err),
	// 	logs.String("debug_message", debugMessage))

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
