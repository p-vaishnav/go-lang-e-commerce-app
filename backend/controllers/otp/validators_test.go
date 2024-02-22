package otp

import (
	"backend-commerce/configs"
	"backend-commerce/constants"
	"backend-commerce/database"
	verifications "backend-commerce/dbops/otp_verifications"
	"backend-commerce/entities"
	"backend-commerce/services/otpsvc"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                           validateOTPAlreadySent                           */
/* -------------------------------------------------------------------------- */
func TestValidateOTPAlreadySent(t *testing.T) {
	configs.LoadConfigs()
	var otpVer entities.OTPVerifications
	var err error

	otpVer.CountryCode = "91"
	otpVer.Mobile = "9834649878"
	otpVer.Status = constants.OTP_STATUS.PENDING
	otpVer.Medium = "sms"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := verifications.Gorm(gormDB)

	otpVer, err = g.CreateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.NotEmpty(t, otpVer.PID)

	var req otpsvc.SendOTPReq

	req.Medium = "sms"
	req.Mobile = "9834649878"

	err = validateOTPAlreadySent(c, req)
	assert.Equal(t, err.Error(), "otp already sent to registered mobile")

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

func TestValidateOTPAlreadySentEmail(t *testing.T) {
	configs.LoadConfigs()
	var otpVer entities.OTPVerifications
	var err error

	otpVer.Email = "vaishnav@gmail.com"
	otpVer.Status = constants.OTP_STATUS.PENDING
	otpVer.Medium = "email"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := verifications.Gorm(gormDB)

	otpVer, err = g.CreateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.NotEmpty(t, otpVer.PID)

	var req otpsvc.SendOTPReq

	req.Email = "vaishnav@gmail.com"
	req.Medium = "email"

	err = validateOTPAlreadySent(c, req)
	assert.Equal(t, err.Error(), "otp already sent to registered email")

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

/* -------------------------------------------------------------------------- */
/*                               validateMedium                               */
/* -------------------------------------------------------------------------- */
func TestValidateMediumSms(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Medium = "sms"
	err = validateMedium(c, req)
	assert.Empty(t, err)
}

func TestValidateMediumEmail(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Medium = "email"
	err = validateMedium(c, req)
	assert.Empty(t, err)
}

func TestValidateMediumWhatsApp(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Medium = "whatsapp"
	err = validateMedium(c, req)
	assert.Empty(t, err)
}

/* -------------------------------------------------------------------------- */
/*                            validateMobileNumber                            */
/* -------------------------------------------------------------------------- */

func TestValidateMobileNumber(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Mobile = "919834649878"
	req.Medium = "sms"

	err = validateMobileNumber(c, req)
	assert.Empty(t, err)
}

func TestValidateMobileNumberInvalid(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Mobile = "919834"
	req.Medium = "sms"

	err = validateMobileNumber(c, req)
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "invalid mobile number provided")
}

func TestValidateMobileNumberMediumEmail(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Mobile = ""
	req.Medium = "email"

	// NOTE: here mobile number won't be checked as medium provided is email
	err = validateMobileNumber(c, req)
	assert.Empty(t, err)
}

/* -------------------------------------------------------------------------- */
/*                                validateEmail                               */
/* -------------------------------------------------------------------------- */
func TestValidateEmail(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Email = "vaishnav@gmail.com"
	req.Medium = "email"

	err = validateEmail(c, req)
	assert.Empty(t, err)
}

func TestValidateEmailInvalid(t *testing.T) {
	var err error
	var req otpsvc.SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req.Email = "vaishnav"
	req.Medium = "email"

	err = validateEmail(c, req)
	assert.NotEmpty(t, err)
	assert.Equal(t, err.Error(), "invalid email provided")
}
