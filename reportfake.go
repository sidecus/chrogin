package main

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/gin-gonic/gin"
)

type FakeDataPoint struct {
	date string
	data int
}

type FakeReportProvider struct {
}

func (p *FakeReportProvider) GetName() string {
	return "fake"
}

func (p *FakeReportProvider) GetTemplate() string {
	return "fake.tmpl"
}

func (p *FakeReportProvider) GetPayload() gin.H {
	data := p.acquireData()
	dataStr := p.formatData(data)

	return gin.H{
		"Title": "fake report",
		"Data":  template.JS(dataStr),
	}
}

func (*FakeReportProvider) acquireData() []FakeDataPoint {
	return []FakeDataPoint{
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
}

func (*FakeReportProvider) formatData(data []FakeDataPoint) string {
	var sb strings.Builder

	sb.WriteString("[")
	for _, point := range data {
		sb.WriteString(fmt.Sprintf("['%s',%d],", template.JSEscapeString(point.date), point.data))
	}
	sb.WriteString("]")

	return sb.String()
}
