package otp

import (
	"backend-commerce/configs"
	"backend-commerce/constants"
	"backend-commerce/database"
	verifications "backend-commerce/dbops/otp_verifications"
	"backend-commerce/entities"
	"backend-commerce/services/otpsvc"
	"fmt"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

/* -------------------------------------------------------------------------- */
/*                               ValidateSendOTP                              */
/* -------------------------------------------------------------------------- */

func validateSendOTPReq(ctx *gin.Context) (otpsvc.SendOTPReq, error) {
	var req otpsvc.SendOTPReq
	var err error

	err = ctx.ShouldBindJSON(&req)
	if err != nil {
		return req, err
	}

	err = validateOTPAlreadySent(ctx, req)
	if err != nil {
		return req, err
	}

	err = validateMedium(ctx, req)
	if err != nil {
		return req, err
	}

	err = validateMobileNumber(ctx, req)
	if err != nil {
		return req, err
	}

	err = validateEmail(ctx, req)
	if err != nil {
		return req, err
	}

	// NOTE: my app is being launched in indian subcontinent only

	return req, err
}

func validateOTPAlreadySent(ctx *gin.Context, req otpsvc.SendOTPReq) error {
	var err error

	gormDB, _ := database.InitDB()
	otpVerGorm := verifications.Gorm(gormDB)

	otpVer, err := otpVerGorm.ListOTPVerification(ctx, req.Mobile, req.Email, constants.OTP_STATUS.PENDING, req.Medium)
	if err != nil {
		return errors.New("query to db failed")
	}

	if otpVer.CreatedAt.Add(time.Duration(configs.Otp.ValidTime) * time.Minute).After(time.Now()) {
		if req.Mobile != "" {
			return errors.New("otp already sent to registered mobile")
		} else {
			return errors.New("otp already sent to registered email")
		}
	}

	return err
}

func validateMedium(ctx *gin.Context, req otpsvc.SendOTPReq) error {
	var err error

	_, exists := constants.OTP_MEDIUM[req.Medium]
	if !exists {
		return errors.New("medium provided is invalid")
	}

	return err
}

func validateMobileNumber(ctx *gin.Context, req otpsvc.SendOTPReq) error {
	var err error

	var smsMed = "sms"
	var otpMed = "otp"
	var checkMobile = false

	for key, _ := range constants.OTP_MEDIUM {
		if (key == smsMed || key == otpMed) && req.Medium == key {
			checkMobile = true
		}
	}

	if !checkMobile {
		return err
	}

	// TODO: correct the regex, study and apply check - Vaishnav
	success, _ := regexp.MatchString(constants.Regex.Mobile, req.Mobile)
	if !success {
		return errors.New("invalid mobile number provided")
	}
	return err
}

func validateEmail(ctx *gin.Context, req otpsvc.SendOTPReq) error {
	var err error

	var emailMed = "email"
	var checkEmail = false

	for key, _ := range constants.OTP_MEDIUM {
		if key == emailMed && req.Medium == key {
			checkEmail = true
		}
	}

	if !checkEmail {
		return err
	}

	// TODO: correct is behaviour
	success, _ := regexp.MatchString(constants.Regex.Email, req.Email)
	if !success {
		return errors.New("invalid email provided")
	}

	return err
}

/* -------------------------------------------------------------------------- */
/*                            validateResendOTPReq                            */
/* -------------------------------------------------------------------------- */

func validateResendOTPReq(ctx *gin.Context) error {
	var err error
	var otpVer entities.OTPVerifications

	otpVer, err = validateVerificationID(ctx)
	if err != nil {
		return err
	}

	err = validateOTPStatus(ctx, otpVer)
	if err != nil {
		return err
	}

	err = validateOTPTimeLimit(ctx, otpVer)
	if err != nil {
		return err
	}

	return err
}

func validateVerificationID(ctx *gin.Context) (entities.OTPVerifications, error) {
	var err error
	var pid string
	var otpVer entities.OTPVerifications

	pid = ctx.Param("pid")
	if pid == "" {
		return otpVer, errors.New("empty pid provided")
	}

	gormDB, _ := database.InitDB()
	otpVerGorm := verifications.Gorm(gormDB)

	otpVer, err = otpVerGorm.ListOTPVerificationByPID(ctx, pid)
	if err != nil {
		return otpVer, errors.New("invalid otp_pid provided")
	}

	return otpVer, err
}

func validateOTPStatus(ctx *gin.Context, otpVer entities.OTPVerifications) error {
	var err error

	if otpVer.Status == constants.OTP_STATUS.FAILURE || otpVer.Status == constants.OTP_STATUS.EXPIRED {
		return errors.New("unprocessable request")
	}

	if otpVer.Status == constants.OTP_STATUS.SUCCESS {
		return errors.New("otp already verified")
	}

	return err
}

func validateOTPTimeLimit(ctx *gin.Context, otpVer entities.OTPVerifications) error {
	var err error

	coolDownTime := otpVer.UpdatedAt.Add(time.Duration(configs.Otp.ValidTime) * time.Second)
	if coolDownTime.After(time.Now()) {
		seconds := coolDownTime.Sub(time.Now()).Seconds()
		return errors.New("please try after," + fmt.Sprintf("%.0f", seconds) + "seconds.")
	}

	return err
}
