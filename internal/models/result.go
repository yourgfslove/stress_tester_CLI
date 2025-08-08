package models

import "time"

type Result struct {
	URL        string
	StatusCode int
	Duration   time.Duration
	Success    bool
	Error      string
}
