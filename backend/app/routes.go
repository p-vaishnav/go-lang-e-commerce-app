package app

import (
	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	var r *gin.Engine

	r = gin.Default()
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "success",
		})
	})

	r.Run()
}
