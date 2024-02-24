package otp

import (
	"backend-commerce/services/otpsvc"
)

func Handler(otpSvc otpsvc.Interface) *otpHandler {
	return &otpHandler{
		otpSvc: otpSvc,
	}
}

type otpHandler struct {
	otpSvc otpsvc.Interface
}
