package authsvc

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
)

func (s *authSvc) VerifyToken(ctx *gin.Context, tokenString string) (string, error) {
	var err error
	var jwtToken *jwt.Token
	var user_pid string

	jwtToken, err = jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// jwt.SigningMethodHS256
		_, ok := token.Method.(*jwt.SigningMethodHMAC) // NOTE: this change
		if !ok {
			return nil, errors.New("unauthorized")
		}

		return "", nil
	})

	if err != nil {
		return user_pid, errors.Wrap(err, "token provided is invalid")
	}

	if !jwtToken.Valid {
		return user_pid, errors.Wrap(err, "[token.Valid] token is not valid")
	}

	claims := jwtToken.Claims.(jwt.MapClaims)
	user_pid = claims["user_pid"].(string)
	return user_pid, err

}
