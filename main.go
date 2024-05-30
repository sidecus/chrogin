package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())
	r.Static("/assets", "./assets")
	// r.StaticFS("/pdf", gin.Dir("./pdf", true))

	r.LoadHTMLGlob("templates/*")

	rg := r.Group("/reports")
	{
		rg.GET("/", reportListHandler)
		rg.GET("/:name", reportHandler)
		rg.GET(fmt.Sprintf("/:name/%s", config.DownloadSuffix), downloadHandler)
	}

	r.Run(fmt.Sprintf(":%d", config.Port))
}
