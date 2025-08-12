package output

import (
	"fmt"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

type Outputer interface {
	Out(summary models.StressSummary) error
}

const (
	typeJson    = "JSON"
	typeTXT     = "TXT"
	typeConsole = "CONSOLE"
)

func NewOutputer(outputType, outputPath string) (Outputer, error) {
	switch outputType {
	case typeConsole:
		return newConsoleOutputer(), nil
	case typeJson:
		return newJSONOutputer(outputPath), nil
	default:
		return nil, fmt.Errorf("wrong output type")
	}
}
