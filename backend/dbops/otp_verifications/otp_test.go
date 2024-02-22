package verifications

import (
	"backend-commerce/configs"
	"backend-commerce/constants"
	"backend-commerce/database"
	"backend-commerce/entities"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                          TestCreateOTPVerification                         */
/* -------------------------------------------------------------------------- */
func TestCreateOTPVerification(t *testing.T) {
	configs.LoadConfigs()
	var otpVer entities.OTPVerifications
	var err error

	otpVer.CountryCode = "91"
	otpVer.Mobile = "9834649878"
	otpVer.Medium = "whatsapp"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := Gorm(gormDB)

	otpVer, err = g.CreateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.NotEmpty(t, otpVer.PID)

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

/* -------------------------------------------------------------------------- */
/*                             ListOTPVerification                            */
/* -------------------------------------------------------------------------- */
func TestFindOTPVerificationMobile(t *testing.T) {
	configs.LoadConfigs()
	var otpVer1 entities.OTPVerifications
	var mobile string
	var err error

	mobile = "9834649111"

	otpVer1.CountryCode = "91"
	otpVer1.Mobile = mobile
	otpVer1.Status = constants.OTP_STATUS.PENDING
	otpVer1.Medium = "whatsapp"
	otpVer1.Provider = "wati"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := Gorm(gormDB)

	otpVer1, err = g.CreateOTPVerification(c, otpVer1)
	assert.Empty(t, err)

	var otpVer2 entities.OTPVerifications

	otpVer2.CountryCode = "44"
	otpVer2.Mobile = mobile
	otpVer2.Status = constants.OTP_STATUS.PENDING
	otpVer2.Medium = "whatsapp"
	otpVer2.Provider = "meta"

	otpVer2, err = g.CreateOTPVerification(c, otpVer2)
	assert.Empty(t, err)

	res, err := g.ListOTPVerification(c, mobile, "", constants.OTP_STATUS.PENDING, "whatsapp")
	assert.Empty(t, err)
	assert.Equal(t, res.Provider, "meta")

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer1.PID).Delete(&otpVer1)
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer2.PID).Delete(&otpVer2)
	})
}

func TestFindOTPVerificationEmail(t *testing.T) {
	configs.LoadConfigs()
	var otpVer1 entities.OTPVerifications
	var email string
	var err error

	email = "vaishnav@zoop.one"

	otpVer1.Email = email
	otpVer1.Status = constants.OTP_STATUS.PENDING
	otpVer1.Medium = "email"
	otpVer1.Provider = "send-grid"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := Gorm(gormDB)

	otpVer1, err = g.CreateOTPVerification(c, otpVer1)
	assert.Empty(t, err)

	var otpVer2 entities.OTPVerifications

	otpVer2.Email = email
	otpVer2.Status = constants.OTP_STATUS.PENDING
	otpVer2.Medium = "email"
	otpVer2.Provider = "ses"

	otpVer2, err = g.CreateOTPVerification(c, otpVer2)
	assert.Empty(t, err)

	res, err := g.ListOTPVerification(c, "", email, constants.OTP_STATUS.PENDING, "email")
	assert.Empty(t, err)
	assert.Equal(t, res.Provider, "ses")

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer1.PID).Delete(&otpVer1)
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer2.PID).Delete(&otpVer2)
	})
}

/* -------------------------------------------------------------------------- */
/*                        TestFindOTPVerificationByPID                        */
/* -------------------------------------------------------------------------- */
func TestFindOTPVerificationByPID(t *testing.T) {
	configs.LoadConfigs()
	var otpVer entities.OTPVerifications
	var err error

	otpVer.CountryCode = "91"
	otpVer.Mobile = "9834649878"
	otpVer.Medium = "whatsapp"

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := Gorm(gormDB)

	otpVer, err = g.CreateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.NotEmpty(t, otpVer.PID)

	res, err := g.ListOTPVerificationByPID(c, otpVer.PID)
	assert.Empty(t, err)
	assert.Equal(t, res.PID, otpVer.PID)

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}

/* -------------------------------------------------------------------------- */
/*                          UpdateOTPVerification                         */
/* -------------------------------------------------------------------------- */
func TestUpdateOTPVerification(t *testing.T) {
	configs.LoadConfigs()
	var otpVer entities.OTPVerifications
	var err error

	otpVer.CountryCode = "91"
	otpVer.Mobile = "9834649878"
	otpVer.Medium = "whatsapp"
	otpVer.Status = constants.OTP_STATUS.PENDING

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	gormDB, _ := database.InitDB()
	g := Gorm(gormDB)

	otpVer, err = g.CreateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.NotEmpty(t, otpVer.PID)

	otpVer.Status = constants.OTP_STATUS.FAILURE
	otpVer, err = g.UpdateOTPVerification(c, otpVer)
	assert.Empty(t, err)
	assert.Equal(t, otpVer.Status, constants.OTP_STATUS.FAILURE)

	t.Cleanup(func() {
		gormDB.Model(&entities.OTPVerifications{}).Where("otp_verifications_pid = ?", otpVer.PID).Delete(&otpVer)
	})
}
