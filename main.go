package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/pdf", gin.Dir("./pdf", true))

	// load HTML templates
	router.LoadHTMLGlob("templates/*")

	m := newReportManager()

	report_router := router.Group("/reports")
	{
		// report lists
		report_router.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, m.providerMap)
		})

		// report handler
		report_router.GET("/:name", reportHandler)

		// download handler
		report_router.GET(fmt.Sprintf("/:name/%s", config.DownloadSuffix), downloadHandler)
	}

	router.Run(fmt.Sprintf(":%d", config.Port))
}
