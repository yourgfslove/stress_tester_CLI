package aggregator

import (
	"time"

	"github.com/yourgfslove/stressTester/internal/models"
)

func Aggregator(summary chan<- models.StressSummary, results <-chan models.Result) {
	var sum models.StressSummary
	sum.StatusCodes = make(map[int]int)
	var total time.Duration
	for v := range results {
		total += v.Duration
		if v.Duration > sum.MaxLatency {
			sum.MaxLatency = v.Duration
		}
		sum.URL = v.URL
		if v.Success {
			sum.Success++
			sum.StatusCodes[v.StatusCode]++
		} else {
			sum.Fail++
		}
	}
	sum.AvgLatency = (total / (time.Duration(sum.Fail) + time.Duration(sum.Success)))
	summary <- sum
	close(summary)
}
