package authsvc

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestVerifyToken(t *testing.T) {
	var err error
	var tokenString string

	authSvc := Handler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tokenString, err = authSvc.CreateToken(c)
	assert.Empty(t, err)
	assert.NotEmpty(t, tokenString)

	user_pid, err := authSvc.VerifyToken(c, tokenString)
	assert.Empty(t, err)
	assert.Equal(t, user_pid, "user_pid")
}
