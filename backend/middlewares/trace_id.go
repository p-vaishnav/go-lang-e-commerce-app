package middlewares

import (
	"backend-commerce/constants"
	"backend-commerce/utils"

	"github.com/gin-gonic/gin"
)

func TraceID() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		traceID := utils.UUID()
		ctx.Request.Header.Add(constants.Headers.TraceID, traceID)
		ctx.Writer.Header().Add(constants.Headers.TraceID, traceID) // adding in the response header
		ctx.Next()
	}
}
