package logs

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// why below code ain't worked with pointer
// type Reader interface {
// 	Read(p []byte) (int, error)
// }

// type Temp struct {
// 	reader Reader
// }

// func (t Temp) Read(p []byte) (int, error) {
// 	return 0, nil
// }

// var t Temp
// fun := func(r Reader) {

// }
// fun(t)

// NOTE: this below function is used in the middleware
func RequestResponseLogs(ctx *gin.Context, gormDB *gorm.DB) {

	// TODO: read about io package in go_lang
	// https://www.youtube.com/watch?v=fzXuXUIewBk&list=PLBRUj12SWlolHbXTpCDjnp08VOj1W94Lc

	ctx.Next() // to go and complete the rest of the work

}
