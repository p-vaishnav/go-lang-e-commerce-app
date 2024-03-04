package otp

import (
	backenderrors "backend-commerce/backend_errors"
	"backend-commerce/utils"

	"github.com/gin-gonic/gin"
)

func (h *otpHandler) SendOTP(ctx *gin.Context) {
	req, err := validateSendOTPReq(ctx)
	if err != nil {
		backenderrors.Validation(ctx, err.Error())
		return
	}

	baseRes, data, err := h.otpSvc.SendOTP(ctx, req)
	if err != nil {
		// NOTE: might need to handle downstream errors as well, and slack notification should be also there
		backenderrors.InternalServer(ctx, err.Error())
		return
	}

	baseRes, err = sendOTPResTransformer(ctx, baseRes, data)
	if err != nil {
		backenderrors.InternalServer(ctx, err.Error())
		return
	}

	utils.ReturnJSONStruct(ctx, baseRes)
}

func (h *otpHandler) ResendOTP(ctx *gin.Context) {
	err := validateResendOTPReq(ctx)
	if err != nil {
		backenderrors.Validation(ctx, err.Error())
		return
	}

	// baseRes, data, err := h.verificationSvc.ResendOTP(ctx)
	// if baseRes.StatusCode != http.StatusOK {
	// 	zooperrors.HandleServiceCodes(ctx, baseRes)
	// 	return
	// }

	// if err != nil {
	// 	zooperrors.InternalServer(ctx, err.Error())
	// 	return
	// }

	// baseRes, err = resendOTPTransformer(ctx, baseRes, data)
	// if err != nil {
	// 	zooperrors.InternalServer(ctx, err.Error())
	// 	return
	// }

	// utils.ReturnJSONStruct(ctx, baseRes)
}

// func (h *verifyHandler) VerifyOTP(ctx *gin.Context) {
// 	req, err := validateVerifyOTPReq(ctx)
// 	if err != nil {
// 		zooperrors.Validation(ctx, err.Error())
// 		return
// 	}

// 	baseRes, data, err := h.verificationSvc.VerifyOTP(ctx, req)
// 	if baseRes.StatusCode != http.StatusOK {
// 		zooperrors.HandleServiceCodes(ctx, baseRes)
// 		return
// 	}

// 	if err != nil {
// 		zooperrors.InternalServer(ctx, err.Error())
// 		return
// 	}

// 	baseRes, err = verifyOTPTransformer(ctx, baseRes, data)
// 	if err != nil {
// 		zooperrors.InternalServer(ctx, err.Error())
// 		return
// 	}

// 	utils.ReturnJSONStruct(ctx, baseRes)
// }
