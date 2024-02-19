package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDBConfigs(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, DB.Host)
	assert.NotEmpty(t, DB.Port)
	assert.NotEmpty(t, DB.Username)
	assert.NotEmpty(t, DB.Database)
	assert.NotEmpty(t, DB.Password)
}

func TestAppConfigs(t *testing.T) {
	LoadConfigs()
	assert.NotEmpty(t, App.Env)
}
