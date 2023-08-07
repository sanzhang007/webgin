package controlers

import (
	"os"

	"github.com/gin-gonic/gin"
)

func Youdou(ctx *gin.Context) {
	b, err := os.ReadFile("static/yudou.txt")
	if err != nil {
		return
	}
	ctx.String(200, string(b))
}
