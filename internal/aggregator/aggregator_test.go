package aggregator

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

func TestAggregator(t *testing.T) {
	input := make(chan models.Result)
	result := make(chan models.StressSummary)
	go Aggregator(result, input)
	results := []models.Result{
		{URL: "someURL", StatusCode: 200, Duration: (time.Second * 3), Success: true},
		{URL: "someURL", StatusCode: 400, Duration: (time.Second * 2), Success: true},
		{URL: "someURL", Duration: (time.Second * 1), Success: false},
		{URL: "someURL", StatusCode: 200, Duration: (time.Second * 3), Success: true},
		{URL: "someURL", StatusCode: 400, Duration: (time.Second * 2), Success: true},
		{URL: "someURL", Duration: (time.Second * 1), Success: false}}
	for _, v := range results {
		input <- v
	}
	close(input)
	sum := <-result
	statusMap := map[int]int{
		400: 2,
		200: 2,
	}
	avgLatency := time.Second * ((3 + 3 + 2 + 2 + 1 + 1) / 6)
	assert.Equal(t, (time.Second * 3), sum.MaxLatency)
	assert.Equal(t, avgLatency, sum.AvgLatency)
	assert.Equal(t, 4, sum.Success)
	assert.Equal(t, 2, sum.Fail)
	assert.Equal(t, statusMap, sum.StatusCodes)

}
