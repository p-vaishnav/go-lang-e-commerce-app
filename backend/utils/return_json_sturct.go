package utils

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

func ReturnJSONStruct(ctx *gin.Context, genericStruct interface{}) {
	var err error

	ctx.Writer.Header().Set("content-type", "application/json")
	err = json.NewEncoder(ctx.Writer).Encode(&genericStruct)
	if err != nil {
		return // does this makes any sense??
	}
}
