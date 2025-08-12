package aggregator

import (
	"time"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

func Aggregator(summary chan<- models.StressSummary, results <-chan models.Result) {
	var sum models.StressSummary
	sum.StatusCodes = make(map[int]int)
	var total time.Duration
	for v := range results {
		total += v.Duration
		if v.Duration > time.Duration(sum.MaxLatency) {
			sum.MaxLatency = models.Duration(v.Duration)
		}
		sum.URL = v.URL
		if v.Success {
			sum.Success++
			sum.StatusCodes[v.StatusCode]++
		} else {
			sum.Fail++
		}
	}
	sum.AvgLatency = models.Duration(total / (time.Duration(sum.Fail) + time.Duration(sum.Success)))
	summary <- sum
	close(summary)
}
