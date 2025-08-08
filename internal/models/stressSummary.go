package models

import "time"

type StressSummary struct {
	URL         string
	Success     int
	Fail        int
	MaxLatency  time.Duration
	AvgLatency  time.Duration
	StatusCodes map[int]int
}
