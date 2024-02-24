package otpsvc

import (
	"backend-commerce/configs"
	"backend-commerce/constants"
	"backend-commerce/database"
	verifications "backend-commerce/dbops/otp_verifications"
	"backend-commerce/entities"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                                  Send OTP                                  */
/* -------------------------------------------------------------------------- */
func TestSendOTPMobile(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var req SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	otpDB := verifications.Gorm(gormDB)

	otpSvc := Handler(otpDB)

	req.Medium = "mobile"
	req.Mobile = "9834649878"

	baseRes, otpVer, err := otpSvc.SendOTP(c, req)
	assert.Empty(t, err)

	assert.Equal(t, baseRes.StatusCode, http.StatusOK)
	assert.Equal(t, baseRes.Status, "success")

	assert.NotEmpty(t, otpVer.PID)
	assert.Equal(t, otpVer.Provider, constants.OTP_PROVIDER.MSG_91)

	t.Cleanup(func() {
		gormDB.Model(entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

func TestSendOTPWhatsapp(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var req SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	otpDB := verifications.Gorm(gormDB)

	otpSvc := Handler(otpDB)

	req.Medium = "whatsapp"
	req.Mobile = "9834649878"

	baseRes, otpVer, err := otpSvc.SendOTP(c, req)
	assert.Empty(t, err)

	assert.Equal(t, baseRes.StatusCode, http.StatusOK)
	assert.Equal(t, baseRes.Status, "success")

	assert.NotEmpty(t, otpVer.PID)
	assert.Equal(t, otpVer.Provider, constants.OTP_PROVIDER.WHATS_APP)

	t.Cleanup(func() {
		gormDB.Model(entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

func TestSendOTPEmail(t *testing.T) {
	configs.LoadConfigs()
	var err error
	var req SendOTPReq

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	otpDB := verifications.Gorm(gormDB)

	otpSvc := Handler(otpDB)

	req.Medium = "email"
	req.Email = "vaishnav@zoop.one"

	baseRes, otpVer, err := otpSvc.SendOTP(c, req)
	assert.Empty(t, err)

	assert.Equal(t, baseRes.StatusCode, http.StatusOK)
	assert.Equal(t, baseRes.Status, "success")

	assert.NotEmpty(t, otpVer.PID)
	assert.Equal(t, otpVer.Provider, constants.OTP_PROVIDER.SEND_GRID)

	t.Cleanup(func() {
		gormDB.Model(entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}
