package controlers

import (
	"fmt"
	"strings"

	"github.com/sanzhang007/webgin/models"

	"github.com/gin-gonic/gin"
)

func Clash(ctx *gin.Context) {
	var nodes []models.Node1
	ctx.Header("Content-Type", "text/plain; charset=utf-8")

	// limit, err := strconv.Atoi(ctx.DefaultQuery("limit", "100"))
	// date := ctx.Query("date")
	// fmt.Printf("err: %v\n", err)
	// avg_i, err := strconv.Atoi(ctx.DefaultQuery("avg_i", "0"))
	// fmt.Printf("err: %v\n", err)

	// avg_j, err := strconv.Atoi(ctx.DefaultQuery("avg_j", "11"))
	// fmt.Printf("err: %v\n", err)

	models.DB.Where(`ping>0 order by ping`).Find(&nodes)
	var builder strings.Builder

	for _, n := range nodes {
		builder.WriteString(fmt.Sprintf("####Ping: %d	AvgSpeed: %.2fMB	MaxSpeed: %.2fMB	SourceUrl: %s	CreateTime: %s	UpdateTime: %s	FailCount: %d\n", n.Ping, float64(n.AvgSpeed)/1024/1024, float64(n.MaxSpeed)/1024/1024, n.Url, n.CreateTime.Format("2006/01/02 15:04"), n.UpdateTime.Format("2006/01/02 15:04"), n.FailCount))
		builder.WriteString(n.Link)
		builder.WriteString("\n")
	}
	clash := templateClash(ClashByte(builder.String()))
	ctx.HTML(200, "config.tmpl", clash)
}
