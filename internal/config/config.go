package config

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	OutputFolder string `yaml:"output_folder" env-required:"true"`
	NumWorkers   int    `yaml:"workers_amount" env-default:"10"`
	OutputType   string `yaml:"output_type"`
}

const (
	outputJSON    = "JSON"
	outputConsole = "CONSOLE"
	outputTXT     = "TXT"
)

var outputTypes = map[string]struct{}{
	outputJSON:    {},
	outputTXT:     {},
	outputConsole: {},
}

func MustLoadConfig() *Config {
	path := fetchcfgPath()
	if path == "" {
		log.Fatal("no config path set")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		panic("config file not found: " + path)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(path, &cfg); err != nil {
		log.Fatalf("can not parse config: %s", err)
	}
	cfg.parsecfg()
	return &cfg
}

func fetchcfgPath() string {
	var path string
	flag.StringVar(&path, "config", "", "set config path")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}
	return path
}

func (cfg *Config) parsecfg() {
	if _, err := os.Stat(cfg.OutputFolder); os.IsNotExist(err) {
		cfg.OutputFolder = ""
		cfg.OutputType = outputConsole
		fmt.Println("No output folder found, only console output")
		return
	}
	if _, ok := outputTypes[strings.ToUpper(cfg.OutputType)]; !ok {
		cfg.OutputType = outputConsole
		fmt.Println("Wrong output type, set CONSOLE")
	}
}
