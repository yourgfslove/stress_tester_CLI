package output

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

const (
	timeFormat = "01_02_2006_15_04_05"
)

type jsonOutput struct {
	path string
}

func newJSONOutputer(outpath string) *jsonOutput {
	return &jsonOutput{
		path: outpath,
	}
}

func (jso jsonOutput) Out(summary models.StressSummary) error {
	preatyURL := urlprety(summary.URL)
	name := fmt.Sprintf("%s/%v-\"%v\".json", jso.path, time.Now().Format(timeFormat), preatyURL)
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

func urlprety(URL string) string {
	URL = strings.TrimPrefix(URL, "http://")
	URL = strings.TrimPrefix(URL, "https://")
	URL = strings.ReplaceAll(URL, "/", ".")
	return URL
}
