package worker

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/yourgfslove/stress_tester_CLI/internal/models"
)

type WorkerPool struct {
	workers int
	Jobs    chan *http.Request
	Results chan models.Result
	client  http.Client
	ctx     context.Context
	cancel  context.CancelFunc
	wg      sync.WaitGroup
}

func NewWorkerPool(workers, buffersize int, client *http.Client) *WorkerPool {
	ctx, cancel := context.WithCancel(context.Background())
	return &WorkerPool{
		workers: workers,
		Jobs:    make(chan *http.Request, buffersize),
		Results: make(chan models.Result, buffersize),
		client:  *client,
		ctx:     ctx,
		cancel:  cancel,
	}
}

func (wp *WorkerPool) Start() {
	for i := 0; i < wp.workers; i++ {
		wp.wg.Add(1)
		go wp.worker(i)
	}
}

func (wp *WorkerPool) worker(id int) {
	defer wp.wg.Done()
	for {
		select {
		case <-wp.ctx.Done():
			return
		case req, ok := <-wp.Jobs:
			if !ok {
				return
			}
			start := time.Now()
			resp, err := wp.client.Do(req)
			var res models.Result
			res.Duration = time.Since(start)
			res.URL = req.URL.String()
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
			wp.Results <- res
		}
	}
}

func (wp *WorkerPool) Stop() {
	close(wp.Jobs)
	wp.cancel()
	wp.wg.Wait()
	close(wp.Results)
}
