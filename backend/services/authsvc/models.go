package authsvc

import "github.com/golang-jwt/jwt"

type UserClaims struct {
	UserPID string `json:"user_pid"`
	jwt.StandardClaims
}
