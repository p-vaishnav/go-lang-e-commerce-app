package authsvc

import (
	"backend-commerce/configs"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (s *authSvc) CreateToken(ctx *gin.Context) (string, error) {
	var accessToken *jwt.Token
	var tokenString string
	var err error

	claims := UserClaims{
		UserPID: "user_pid",
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(time.Duration(configs.Token.AccessExpiryTime) * time.Minute).Unix(),
		},
	}

	accessToken = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err = accessToken.SignedString([]byte(configs.Token.AccessSecret))
	if err != nil {
		return tokenString, errors.Wrap(err, "[CreateToken][SignedString]")
	}

	return tokenString, err
}
