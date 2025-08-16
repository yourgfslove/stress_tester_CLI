package output

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

const (
	timeFormat = "01_02_2006_15_04_05"
)

// represents an output type for JSON files
type jsonOutput struct {
	path string
}

// constructor for jsonOutput
func newJSONOutputer(outpath string) *jsonOutput {
	return &jsonOutput{
		path: outpath,
	}
}

// Out is creating a new JSON file of stress test summary
func (jso jsonOutput) Out(summary models.StressSummary) error {
	name := namegenerator(jso.path, summary.URL, "json")
	outFile, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("error to create file: %v", err)
	}
	defer outFile.Close()
	encoder := json.NewEncoder(outFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(summary); err != nil {
		return err
	}
	return nil
}
