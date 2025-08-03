package worker

import (
	"context"
	"net/http"
	"time"
)

type Result struct {
	StatusCode int
	Duration   time.Duration
	Success    bool
	Error      string
}

func Run(ctx context.Context, Client *http.Client, results chan<- Result, jobs <-chan *http.Request) {
	for {
		select {
		case <-ctx.Done():
			return
		case req := <-jobs:
			start := time.Now()
			resp, err := Client.Do(req)
			var res Result
			res.Duration = time.Since(start)

			if err != nil {
				res.Success = false
				res.Error = err.Error()
			} else {
				res.Success = true
				res.StatusCode = resp.StatusCode
				resp.Body.Close()
			}
			results <- res
		}
	}
}
