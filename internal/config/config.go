package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	OutputFolder string `yaml:"output_folder" env-required:"true"`
	NumWorkers   int    `yaml:"workers_amount" env-default:"10"`
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
