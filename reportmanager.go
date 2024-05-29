package main

type ReportManager struct {
	providers   []IReportProvider
	providerMap map[string]IReportProvider
}

func newReportManager() *ReportManager {
	m := &ReportManager{
		providers: []IReportProvider{
			&FakeReportProvider{},
		},
	}

	m.providerMap = make(map[string]IReportProvider, len(m.providers))
	for _, p := range m.providers {
		m.providerMap[p.GetName()] = p
	}

	return m
}
