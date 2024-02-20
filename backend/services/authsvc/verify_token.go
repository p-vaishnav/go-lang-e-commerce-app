package authsvc

import (
	"backend-commerce/configs"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (s *authSvc) VerifyToken(ctx *gin.Context, accessToken string) (string, error) {
	var err error
	var parsedAccessToken *jwt.Token
	var userPID string

	parsedAccessToken, err = jwt.ParseWithClaims(accessToken, &UserClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.Token.AccessSecret), nil
	})

	if err != nil {
		return userPID, errors.Wrap(err, "[ParseWithClaims]")
	}

	claims := parsedAccessToken.Claims.(*UserClaims)

	return claims.UserPID, err

}
