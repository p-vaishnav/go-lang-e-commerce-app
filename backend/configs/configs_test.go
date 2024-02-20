package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/* -------------------------------------------------------------------------- */
/*                                loadOTPConfig                               */
/* -------------------------------------------------------------------------- */
func TestOTPConfigs(t *testing.T) {
	LoadConfigs()

	assert.NotEmpty(t, Otp.Count)
	assert.NotEmpty(t, Otp.ValidTime)
	assert.NotEmpty(t, Otp.VerifyLimit)
}

/* -------------------------------------------------------------------------- */
/*                               loadTokenConfig                              */
/* -------------------------------------------------------------------------- */
func TestTokenConfigs(t *testing.T) {
	LoadConfigs()

	assert.NotEmpty(t, Token.AccessSecret)
	assert.NotEmpty(t, Token.AccessExpiryTime)
	assert.NotEmpty(t, Token.RefreshSecret)
	assert.NotEmpty(t, Token.RefreshExpiryTime)
}

/* -------------------------------------------------------------------------- */
/*                                loadDBConfig                                */
/* -------------------------------------------------------------------------- */
func TestDBConfigs(t *testing.T) {
	LoadConfigs()

	assert.NotEmpty(t, DB.Host)
	assert.NotEmpty(t, DB.Port)
	assert.NotEmpty(t, DB.Database)
	assert.NotEmpty(t, DB.Username)
	assert.NotEmpty(t, DB.Password)
}
