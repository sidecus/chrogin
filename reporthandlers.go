package main

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"

	"github.com/gin-gonic/gin"
)

// report list handler
func reportListHandler(c *gin.Context) {
	list := m.GetSummary()
	logger.Info("Returning built-in reportlist.", "num", len(list))

	c.JSON(http.StatusOK, list)
}

// embeded report handler
type ReportReq struct {
	Name string `uri:"name" binding:"required"`
}

func reportHandler(c *gin.Context) {
	var req ReportReq
	if err := c.ShouldBindUri(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid report request"})
		return
	}

	logger.Info("Rendering built-in report", "name", req.Name)

	p, err := m.GetProvider(req.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	c.HTML(http.StatusOK, p.GetTemplate(), p.GetPayload())
}

// download handler
type DownloadReq struct {
	Name string `json:"name"` // embeded report name
	Uri  string `json:"uri"`  // or, a report uri
}

func getReportUri(req DownloadReq) (string, error) {
	if req.Name == "" && req.Uri == "" {
		return "", fmt.Errorf("name and uri cannot be empty at the same time")
	}

	if req.Name != "" {
		if _, err := m.GetProvider(req.Name); err != nil {
			return "", err
		}

		// Try to generate the report uri based on its name
		return fmt.Sprintf(EmbededReportUriFmtString, config.Port, req.Name), nil
	}

	_, err := url.Parse(req.Uri)
	if err != nil {
		return "", err
	}

	return req.Uri, nil
}

func getTempFileName() (string, error) {
	f, err := os.CreateTemp("", "*.pdf")
	if err != nil {
		return "", err
	}

	// We need to close the file
	f.Close()
	return f.Name(), nil
}

func downloadHandler(c *gin.Context) {
	var req DownloadReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	// Validate and get report uri
	uri, err := getReportUri(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	logger.Info("Generating report from uri", "uri", uri)

	// get temp file name
	tempFile, err := getTempFileName()
	if err != nil {
		logger.Error("Failed to create temp file.", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "create report file failed"})
		return
	}
	defer os.Remove(tempFile)

	// Run report generation
	cmd := exec.Command(
		"chromium",
		"--no-sandbox",
		"--headless",
		"-disable-gpu",
		"--no-pdf-header-footer",
		fmt.Sprintf("--print-to-pdf=%s", tempFile),
		fmt.Sprintf("--timeout=%d", config.ChromiumTimeoutSeconds),
		uri,
	)
	if err = cmd.Run(); err != nil {
		logger.Error("Report generation failed.", "error", err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "report generation failed"})
		return
	}

	// Return the file
	c.File(tempFile)
}
