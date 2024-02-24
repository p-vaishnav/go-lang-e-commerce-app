package otpsvc

import (
	"backend-commerce/constants"
	"backend-commerce/entities"
	"backend-commerce/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func (s *otpSvc) SendOTP(ctx *gin.Context, req SendOTPReq) (models.BaseResponse, entities.OTPVerifications, error) {
	var err error
	var baseRes models.BaseResponse
	var otpVer entities.OTPVerifications
	var otp string
	var provider string

	baseRes.StatusCode = http.StatusInternalServerError
	baseRes.Status = "error"
	baseRes.Message = "internal server error"

	// otp = utils.GenerateOTP(ctx)
	otp = "123456" // hard-coding the otp as of now

	// TODO: complete the providers
	switch req.Medium {
	case "email":
		provider = constants.OTP_PROVIDER.SEND_GRID
	case "whatsapp":
		provider = constants.OTP_PROVIDER.WHATS_APP
	case "mobile":
		provider = constants.OTP_PROVIDER.MSG_91
	}

	// make db insertion
	otpVer.Email = req.Email
	otpVer.Mobile = req.Mobile
	otpVer.Medium = req.Medium
	otpVer.OTPCount = 1
	otpVer.Provider = provider
	otpVer.Status = constants.OTP_STATUS.PENDING
	otpVer.OTP = otp

	otpVer, err = s.gormDB.CreateOTPVerification(ctx, otpVer)
	if err != nil {
		return baseRes, otpVer, errors.Wrap(err, "[SendOTP]")
	}

	baseRes.StatusCode = http.StatusOK
	baseRes.Status = "success"
	baseRes.Message = "otp sent successfully"

	return baseRes, otpVer, err
}
