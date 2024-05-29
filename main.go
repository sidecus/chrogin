package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func formatData(data []DataPoint) string {
	var sb strings.Builder
	sb.WriteString("[")

	for _, point := range data {
		sb.WriteString(fmt.Sprintf("['%s',%d],", template.JSEscapeString(point.date), point.data))
	}

	sb.WriteString("]")
	return sb.String()
}

func main() {
	router := gin.Default()
	router.Static("/assets", "./assets")
	router.StaticFS("/pdf", gin.Dir("./pdf", true))

	// load HTML templates
	router.LoadHTMLGlob("templates/**/*")

	// about handler
	router.GET("/about", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about/index.tmpl", gin.H{
			"title": "About",
		})
	})

	// report handler
	router.GET("/report", func(c *gin.Context) {
		data := acquireData()
		dataStr := formatData(data)

		c.HTML(http.StatusOK, "report/index.tmpl", gin.H{
			"title": "Users",
			"Data":  template.JS(dataStr),
		})
	})

	router.Run(":8080")
}
