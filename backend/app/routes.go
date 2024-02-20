package app

import (
	"fmt"

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

	v1Group := r.Group("v1")
	{
		otpGroup := v1Group.Group("otp")
		{
			otpGroup.POST("send_otp", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "send-otp-success",
				})
			})
			otpGroup.POST("resend_otp", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "resend-otp-success",
				})
			})
			otpGroup.POST("verify_otp", func(c *gin.Context) {
				c.JSON(200, gin.H{
					"message": "verify-otp-success",
				})
			})
		}
	}

	fmt.Println("initialized all the routes")
	r.Run()
}