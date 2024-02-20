package configs

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

type OtpConfig struct {
	Count       int `split_words:"true" json:"OTP_COUNT"`
	ValidTime   int `split_words:"true" json:"OTP_VALID_TIME"`
	VerifyLimit int `split_words:"true" json:"OTP_VERIFY_LIMIT"`
}

var Otp *OtpConfig

func loadOTPConfig() {
	Otp = &OtpConfig{}

	err := envconfig.Process("OTP", Otp)
	if err != nil {
		log.Fatal(err.Error())
	}
}
