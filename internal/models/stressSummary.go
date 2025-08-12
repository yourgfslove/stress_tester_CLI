package models

type StressSummary struct {
	URL         string      `json:"request_url"`
	Success     int         `json:"success_request"`
	Fail        int         `json:"failed_request"`
	MaxLatency  Duration    `json:"max_latency"`
	AvgLatency  Duration    `json:"average_latency"`
	StatusCodes map[int]int `json:"status_codes"`
}
