package authsvc

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (s *authSvc) CreateToken(ctx *gin.Context) (string, error) {
	var token *jwt.Token
	var tokenString string
	var err error

	var sampleSecretKey = []byte("SecretYouShouldHide")
	fmt.Println(sampleSecretKey)

	token = jwt.New(jwt.SigningMethodEdDSA) // TODO: it didn't worked with it EdDSA

	// claims
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(time.Minute * 10)
	claims["user_pid"] = "user_pid"

	tokenString, err = token.SignedString(sampleSecretKey)
	if err != nil {
		return tokenString, errors.Wrap(err, "[CreateToken]")
	}

	return tokenString, err
}
