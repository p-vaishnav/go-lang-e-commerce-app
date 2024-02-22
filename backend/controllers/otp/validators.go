package otp

import (
	"backend-commerce/configs"
	"backend-commerce/constants"
	"backend-commerce/database"
	verifications "backend-commerce/dbops/otp_verifications"
	"backend-commerce/services/otpsvc"
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
		return req, errors.New("[ValidateSendOTPReq]error while unmarshalling")
	}

	err = validateOTPAlreadySent(ctx, req)
	if err != nil {
		return req, errors.Wrap(err, "[ValidateSendOTPReq]")
	}

	err = validateMedium(ctx, req)
	if err != nil {
		return req, errors.Wrap(err, "[ValidateSendOTPReq]")
	}

	err = validateMobileNumber(ctx, req)
	if err != nil {
		return req, errors.Wrap(err, "[validateSendOTPReq]")
	}

	err = validateEmail(ctx, req)
	if err != nil {
		return req, errors.Wrap(err, "[validateSendOTPReq]")
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
