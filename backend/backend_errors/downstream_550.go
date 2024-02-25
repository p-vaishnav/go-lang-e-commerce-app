package backenderrors

import (
	"backend-commerce/models"

	"github.com/gin-gonic/gin"
)

func Downstream(ctx *gin.Context, err string) {
	var res models.BaseResponse
	errorCode := 550

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = err
	res.Status = "error"

	// logs.Warn("DOWNSTREAM",
	// 	logs.Int("status_code", errorCode),
	// 	logs.String("status", res.Status),
	// 	logs.String("message", err))

	// TODO: generate some kind of mail notification service where developer is notified via mail or phone

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
