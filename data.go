package main

type DataPoint struct {
	date string
	data int
}

func acquireData() []DataPoint {
	return []DataPoint{
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
