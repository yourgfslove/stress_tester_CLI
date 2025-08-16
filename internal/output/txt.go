package output

import (
	"fmt"
	"os"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

// represents an output type for TXT files
type txtOutput struct {
	path string
}

// constructor for txtOutput
func newTXTOutputer(outpath string) *txtOutput {
	return &txtOutput{
		path: outpath,
	}
}

// Out is creating a new TXT file of stress test summary
func (txto txtOutput) Out(summary models.StressSummary) error {
	name := namegenerator(txto.path, summary.URL, "txt")
	outFile, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("error to create file: %v", err)
	}
	defer outFile.Close()
	fmt.Fprintf(outFile, "Number of success requests: %d\n", summary.Success)
	fmt.Fprintf(outFile, "Number of failed requests: %d\n", summary.Fail)
	fmt.Fprintf(outFile, "Average latency: %v\n", summary.AvgLatency)
	fmt.Fprintf(outFile, "Max latency: %v\n", summary.MaxLatency)
	fmt.Fprintln(outFile, "Status Codes:")
	for k, v := range summary.StatusCodes {
		fmt.Fprintf(outFile, " -%d: %d\n", k, v)
	}
	return nil
}
