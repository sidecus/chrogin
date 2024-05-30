package main

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

type FakeReportProvider struct {
	Name     string
	Template string
}

func newFakeReportProvider() *FakeReportProvider {
	p := &FakeReportProvider{}
	p.Name = "fake"
	p.Template = "fake.tmpl"

	return p
}

func (p *FakeReportProvider) GetName() string {
	return p.Name
}

func (p *FakeReportProvider) GetTemplate() string {
	return p.Template
}

func (p *FakeReportProvider) GetPayload() gin.H {
	data := []FakeDataPoint{
		{date: "2019-10-10", data: 200},
		{date: "2019-10-11", data: 560},
		{date: "2019-10-12", data: 750},
		{date: "2019-10-13", data: 580},
		{date: "2019-10-14", data: 250},
		{date: "2019-10-15", data: 300},
		{date: "2019-10-16", data: 450},
		{date: "2019-10-17", data: 300},
		{date: "2019-10-18", data: 100},
	}

	// Convert data to JS array
	var sb strings.Builder

	sb.WriteString("[")
	for _, point := range data {
		sb.WriteString(fmt.Sprintf("['%s',%d],", template.JSEscapeString(point.date), point.data))
	}
	sb.WriteString("]")

	dataStr := sb.String()

	// Construct payload
	return gin.H{
		"Title": "fake report",
		"Data":  template.JS(dataStr),
	}
}

type FakeDataPoint struct {
	date string
	data int
}
