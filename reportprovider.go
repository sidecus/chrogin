package main

import "github.com/gin-gonic/gin"

type IReportProvider interface {
	GetName() string
	GetTemplate() string
	GetPayload() gin.H
}
