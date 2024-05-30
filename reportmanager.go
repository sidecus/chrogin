package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IReportProvider interface {
	GetName() string
	GetTemplate() string
	GetPayload() gin.H
}

type ReportManager struct {
	Providers   map[string]IReportProvider
	ProviderMap map[string]string
}

func (m *ReportManager) GetSummary() map[string]string {
	return m.ProviderMap
}

func (m *ReportManager) GetProvider(name string) (IReportProvider, error) {
	p, ok := m.Providers[name]
	if !ok {
		return nil, fmt.Errorf("report %s doesn't exist", name)
	}

	return p, nil
}

func newReportManager() *ReportManager {
	m := &ReportManager{}
	m.Providers = make(map[string]IReportProvider, len(m.Providers))
	m.ProviderMap = make(map[string]string, len(m.Providers))

	providers := []IReportProvider{
		newFakeReportProvider(),
	}

	for _, p := range providers {
		m.Providers[p.GetName()] = p
		m.ProviderMap[p.GetName()] = p.GetTemplate()
	}

	return m
}
