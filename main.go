package main

import (
	"webgin/controlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	//
	r.LoadHTMLGlob("template/*")
	r.StaticFile("static", "./static")

	r.GET("/nodes", controlers.Nodes)
	// r.GET("/nodes/:avg", controlers.Nodes)
	// r.GET("/nodes/:avg/:limit", controlers.Nodes)
	// r.GET("/nodes/:avg/:limit/:date", controlers.Nodes)

	r.GET("/clash/", controlers.Clash)
	// r.GET("/clash/:avg", controlers.Clash)
	// r.GET("/clash/:avg/:limit/", controlers.Clash)
	// r.GET("/clash/:avg/:limit/:date", controlers.Clash)

	r.GET("/clash-convert/", controlers.ClashConvert)
	r.GET("/nodes-convert")
	r.Run(":80")
}
