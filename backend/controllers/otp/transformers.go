package otp

import (
	"backend-commerce/entities"
	"backend-commerce/models"

	"github.com/gin-gonic/gin"
)

func sendOTPResTransformer(ctx *gin.Context, baseRes models.BaseResponse, data entities.OTPVerifications) (models.BaseResponse, error) {
	var err error
	return baseRes, err
}
