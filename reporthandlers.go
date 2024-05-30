package main

import (
	"fmt"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/gin-gonic/gin"
)

type Req struct {
	Name string `uri:"name" binding:"required"`
}

var m = newReportManager()

func reportListHandler(c *gin.Context) {
	c.JSON(http.StatusOK, m.GetSummary())
}

func reportHandler(c *gin.Context) {
	var req Req
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid report request"})
		return
	}

	p, err := m.GetProvider(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.HTML(http.StatusOK, p.GetTemplate(), p.GetPayload())
}

func downloadHandler(c *gin.Context) {
	f, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "create report file failed"})
		return
	}

	// We need to close the file
	f.Close()
	defer os.Remove(f.Name())

	// Try to generate the report using the given file
	relativeUrl := strings.TrimSuffix(c.Request.URL.String(), config.DownloadSuffix)
	url := fmt.Sprintf("http://localhost:%d%s", config.Port, relativeUrl)
	cmd := exec.Command(
		"chromium",
		"--no-sandbox",
		"--headless",
		"-disable-gpu",
		"--no-pdf-header-footer",
		fmt.Sprintf("--print-to-pdf=%s", f.Name()),
		"--timeout=5000",
		url,
	)

	if err = cmd.Run(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "report generation failed"})
		return
	}

	// Return the file
	c.File(f.Name())
}
