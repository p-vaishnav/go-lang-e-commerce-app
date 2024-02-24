package utils

import (
	"fmt"
	"math/rand"

	"github.com/gin-gonic/gin"
)

func GenerateOTP(ctx *gin.Context) string {
	var result string

	result = ""
	for i := 0; i < 6; i++ {
		result = result + fmt.Sprintf("%d", rand.Intn(10))
	}

	return result
}
