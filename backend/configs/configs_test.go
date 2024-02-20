package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOTPConfigs(t *testing.T) {
	LoadConfigs()

	assert.NotEmpty(t, Otp.Count)
	assert.NotEmpty(t, Otp.ValidTime)
	assert.NotEmpty(t, Otp.VerifyLimit)
}
