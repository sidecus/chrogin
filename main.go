package main

import (
	"fmt"
	"log/slog"

	"github.com/gin-gonic/gin"
)

var config = newConfig()
var m = newReportManager()
var logger = slog.Default()

func main() {
	r := gin.Default()

	r.Use(gin.Recovery())

	r.Static("/assets", "./assets")
	r.LoadHTMLGlob("templates/*")

	rg := r.Group("/reports")
	{
		rg.GET("/", reportListHandler)
		rg.GET("/:name", reportHandler)
	}
	r.POST("/download", downloadHandler)

	logger.Info("Starting server on port", "port", config.Port)
	r.Run(fmt.Sprintf(":%d", config.Port))
}
