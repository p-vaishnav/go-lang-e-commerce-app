package authsvc

import (
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestCreateToken(t *testing.T) {
	var err error
	var tokenString string

	authSvc := Handler()

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	tokenString, err = authSvc.CreateToken(c)
	assert.Empty(t, err)
	assert.NotEmpty(t, tokenString)
}
