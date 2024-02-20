package middlewares

import (
	"github.com/gin-gonic/gin"
)

func RequestResponseLogs() gin.HandlerFunc {
	// TODO: here a insert query should be made on request response logs
	return func(ctx *gin.Context) {

		ctx.Next()
	}
}
