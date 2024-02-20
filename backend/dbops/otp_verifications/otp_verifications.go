package verifications

import (
	"backend-commerce/constants"
	"backend-commerce/entities"
	"backend-commerce/utils"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	CreateOTPVerification(*gin.Context, entities.OTPVerifications) (entities.OTPVerifications, error)
	DeleteOTPVerification(*gin.Context, entities.OTPVerifications) error
	FindOTPVerification(*gin.Context, string, string, string, string) (entities.OTPVerifications, error)
	FindOTPVerificationByPID(*gin.Context, string) (entities.OTPVerifications, error)
	UpdateOTPVerification(*gin.Context, entities.OTPVerifications) (entities.OTPVerifications, error)
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
type otpVerifications struct {
	DB *gorm.DB
}

func Gorm(DB *gorm.DB) *otpVerifications {
	return &otpVerifications{
		DB: DB,
	}
}

/* -------------------------------------------------------------------------- */
/*                                   Methods                                  */
/* -------------------------------------------------------------------------- */

/* -------------------------------------------------------------------------- */
/*                            CreateOTPVerification                           */
/* -------------------------------------------------------------------------- */
func (g *otpVerifications) CreateOTPVerification(ctx *gin.Context, otpVer entities.OTPVerifications) (entities.OTPVerifications, error) {
	var err error

	otpVer.PID = utils.UUIDWithPrefix(constants.Prefix.OTP_VERIFICATION)

	db := g.DB.Session(&gorm.Session{}).Create(&otpVer)
	err = db.Error
	if err != nil {
		return otpVer, errors.Wrap(err, "[CreateOTPVerification]")
	}

	return otpVer, err
}

/* -------------------------------------------------------------------------- */
/*                            DeleteOTPVerification                           */
/* -------------------------------------------------------------------------- */
func (g *otpVerifications) DeleteOTPVerification(ctx *gin.Context, otpVer entities.OTPVerifications) error {
	var err error

	db := g.DB.Session(&gorm.Session{}).Delete(&otpVer)
	err = db.Error
	if err != nil {
		return errors.Wrap(err, "[DeleteOTPVerification]")
	}

	return err
}

/* -------------------------------------------------------------------------- */
/*                             FindOTPVerification                            */
/* -------------------------------------------------------------------------- */
func (g *otpVerifications) FindOTPVerification(ctx *gin.Context, mobile string, email string, status string, medium string) (entities.OTPVerifications, error) {
	var res entities.OTPVerifications
	var err error

	db := g.DB.Session(&gorm.Session{})
	if mobile != "" {
		db.Where("mobile = ? AND status = ? AND medium = ?", mobile, status, medium).Order("created_at desc").Take(&res)
	}

	if email != "" {
		db.Where("email = ? AND status = ? AND medium = ?", email, status, medium).Order("created_at desc").Take(&res)
	}

	return res, err
}

/* -------------------------------------------------------------------------- */
/*                          FindOTPVerificationByPID                          */
/* -------------------------------------------------------------------------- */
func (g *otpVerifications) FindOTPVerificationByPID(ctx *gin.Context, pid string) (entities.OTPVerifications, error) {
	var res entities.OTPVerifications
	var err error

	// TODO: Vaishnav, pid should not be pased explicitly scopes shouldbe there
	db := g.DB.Session(&gorm.Session{})
	err = db.Where("otp_verifications_pid = ?", pid).Take(&res).Error
	if err != nil {
		return res, errors.Wrap(err, "[FindOTPVerificationByPID]")
	}

	return res, err
}

/* -------------------------------------------------------------------------- */
/*                            UpdateOTPVerification                           */
/* -------------------------------------------------------------------------- */
func (g *otpVerifications) UpdateOTPVerification(ctx *gin.Context, otpVer entities.OTPVerifications) (entities.OTPVerifications, error) {
	var err error

	db := g.DB.Session(&gorm.Session{})
	result := db.Where("otp_verifications_pid = ?", otpVer.PID).Updates(&otpVer)

	err = result.Error
	if err != nil {
		return otpVer, errors.Wrap(err, "[UpdateOTPVerification]")
	}

	return otpVer, err
}
