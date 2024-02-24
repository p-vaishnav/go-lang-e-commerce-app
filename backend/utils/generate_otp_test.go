package utils

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGenerateOTP(t *testing.T) {

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	res := GenerateOTP(c)
	assert.NotEmpty(t, res)
	assert.Equal(t, len(res), 6)
}
