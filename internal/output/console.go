package output

import (
	"fmt"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

type consoleOutput struct{}

func newConsoleOutputer() consoleOutput {
	return consoleOutput{}
}

func (co consoleOutput) Out(summary models.StressSummary) error {
	fmt.Printf("Number of success requests: %d\n", summary.Success)
	fmt.Printf("Number of failed requests: %d\n", summary.Fail)
	fmt.Printf("Average latency: %v\n", summary.AvgLatency)
	fmt.Printf("Max latency: %v\n", summary.MaxLatency)
	fmt.Println("Status Codes:")
	for k, v := range summary.StatusCodes {
		fmt.Printf(" -%d: %d\n", k, v)
	}
	return nil
}
