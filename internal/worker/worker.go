package worker

import (
	"context"
	"net/http"
	"time"
)

type Result struct {
	URL        string
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
		case req, ok := <-jobs:
			if !ok {
				return
			}
			start := time.Now()
			resp, err := Client.Do(req)
			var res Result
			res.Duration = time.Since(start)
			res.URL = req.RequestURI
			if err != nil {
				res.Success = false
				res.Error = err.Error()
			} else {
				res.Success = true
				if resp != nil {
					res.StatusCode = resp.StatusCode
					resp.Body.Close()
				}
			}
			results <- res
		}
	}
}
