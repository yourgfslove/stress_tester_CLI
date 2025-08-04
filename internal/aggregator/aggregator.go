package aggregator

import (
	"time"

	"github.com/yourgfslove/stressTester/internal/worker"
)

type StressSummary struct {
	Success     int
	Fail        int
	MaxLatency  time.Duration
	AvgLatency  time.Duration
	StatusCodes map[int]int
}

func Aggregator(summary chan<- StressSummary, results <-chan worker.Result) {
	var sum StressSummary
	sum.StatusCodes = make(map[int]int)
	var total time.Duration
	for v := range results {
		total += v.Duration
		if v.Duration > sum.MaxLatency {
			sum.MaxLatency = v.Duration
		}
		if v.Success {
			sum.Success++
		} else {
			sum.Fail++
		}
		sum.StatusCodes[v.StatusCode]++
	}
	sum.AvgLatency = (total / (time.Duration(sum.Fail) + time.Duration(sum.Success)))
	summary <- sum
	close(summary)
}
