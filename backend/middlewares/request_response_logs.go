package middlewares

import (
	"backend-commerce/database"
	"backend-commerce/logs"

	"github.com/gin-gonic/gin"
)

func RequestResponseLogs() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// TODO: add a config check wether to activate req, res logs or not??
		gormDB, _ := database.InitDB()
		logs.RequestResponseLogs(ctx, gormDB)
		ctx.Next()
	}
}
