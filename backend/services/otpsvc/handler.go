package otpsvc

import (
	verifications "backend-commerce/dbops/otp_verifications"
	"backend-commerce/entities"
	"backend-commerce/models"

	"github.com/gin-gonic/gin"
)

/* -------------------------------------------------------------------------- */
/*                                  Interface                                 */
/* -------------------------------------------------------------------------- */
type Interface interface {
	SendOTP(ctx *gin.Context, req SendOTPReq) (models.BaseResponse, entities.OTPVerifications, error)
}

/* -------------------------------------------------------------------------- */
/*                                   otpSvc                                   */
/* -------------------------------------------------------------------------- */
type otpSvc struct {
	gormDB verifications.Interface
}

/* -------------------------------------------------------------------------- */
/*                                   Handler                                  */
/* -------------------------------------------------------------------------- */
func Handler(gormDB verifications.Interface) *otpSvc {
	return &otpSvc{
		gormDB: gormDB,
	}
}
