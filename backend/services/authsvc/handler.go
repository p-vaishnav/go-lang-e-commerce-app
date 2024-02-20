package authsvc

import "github.com/gin-gonic/gin"

// Redis will store user_pid : true with a time to live of 30 days
// https://pascalallen.medium.com/jwt-authentication-with-go-242215a9b4f8 - "referred from here"

type Interface interface {
	CreateToken(ctx *gin.Context) (string, error)
	VerifyToken(ctx *gin.Context) (string, error)
	DeleteToken(ctx *gin.Context) error
}

type authSvc struct{}

func Handler() *authSvc {
	return &authSvc{}
}
