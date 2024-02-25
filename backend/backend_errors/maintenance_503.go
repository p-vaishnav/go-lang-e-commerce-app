package backenderrors

import (
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// only used for downtime
func Maintenance(ctx *gin.Context) {
	var res models.BaseResponse
	errorCode := http.StatusServiceUnavailable

	res.StatusCode = errorCode
	res.Data = nil
	res.Message = "service down for maintenance. Please try after sometime."
	res.Status = "maintenance"

	ctx.JSON(errorCode, res)
	ctx.Abort()
}
