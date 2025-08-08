package commands

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/pingcap/log"
	"github.com/yourgfslove/stressTester/internal/aggregator"
	"github.com/yourgfslove/stressTester/internal/config"
	"github.com/yourgfslove/stressTester/internal/lib/validation"
	"github.com/yourgfslove/stressTester/internal/models"
	"github.com/yourgfslove/stressTester/internal/worker"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Callback    func(ctx context.Context, args []string) error
}

type Commands map[string]Command

var cmds map[string]Command

func MustInitCommands(MainConf config.Config) Commands {
	cmds = map[string]Command{
		"exit": {
			Name:        "Exit",
			Description: "exiting from CLI",
			Callback:    exit,
		},
		"stress": {
			Name:        "Stress",
			Description: "starting stress test",
			Usage:       " --rps - amount of request per second \n --link - link on stress site \n --method[GET/POST/PATCH] - method",
			Callback: func(ctx context.Context, args []string) error {
				return stressTest(ctx, MainConf, args)
			},
		},
		"help": {
			Name:        "Help",
			Description: "describing all commannds",
			Callback:    help,
		},
	}
	return cmds
}

func exit(ctx context.Context, args []string) error {
	os.Exit(0)
	return nil
}

func stressTest(ctx context.Context, cfg config.Config, args []string) error {
	fs := flag.NewFlagSet("stress", flag.ContinueOnError)

	var rps int
	var link string
	var method string
	var duration int
	var data string
	var contentType string

	fs.IntVar(&rps, "rps", 10, "amount of request per second")
	fs.StringVar(&link, "link", "https://httpbin.org/get", "link on stress site")
	fs.StringVar(&method, "method", "GET", "Request method")
	fs.IntVar(&duration, "sec", 1, "num of seconds doing requests")
	fs.StringVar(&data, "data", "", "data for response")
	fs.StringVar(&contentType, "content/type", "", "content/type header for request")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse parameters: %s", err.Error())
	}

	method = strings.ToUpper(method)

	if err := validation.StressParamsValidate(link, method); err != nil {
		return fmt.Errorf("bad params: %s", err.Error())
	}
	transport := &http.Transport{
		MaxIdleConns:        cfg.NumWorkers + 5,
		MaxIdleConnsPerHost: cfg.NumWorkers + 5,
		MaxConnsPerHost:     cfg.NumWorkers + 5,
	}
	client := http.Client{
		Transport: transport,
		Timeout:   10 * time.Second,
	}

	wp := worker.NewWorkerPool(cfg.NumWorkers, cfg.NumWorkers*3, &client)
	wp.Start()
	timeoutCtx, cancel := context.WithTimeout(ctx, time.Duration(duration)*time.Second)
	defer cancel()

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rps))
		defer ticker.Stop()
		for {
			select {
			case <-timeoutCtx.Done():
				wp.Stop()
				return
			case <-ticker.C:
				req, err := http.NewRequest(method, link, bytes.NewReader([]byte(data)))
				if err != nil {
					log.Error("can not create request")
				}
				if contentType != "" {
					req.Header.Set("Content-Type", contentType)
				}
				req = req.WithContext(timeoutCtx)
				wp.Jobs <- req
			}
		}
	}()

	summary := make(chan models.StressSummary)
	go aggregator.Aggregator(summary, wp.Results)

	sum := <-summary
	fmt.Printf("Number of success requests: %d\n", sum.Success)
	fmt.Printf("Number of failed requests: %d\n", sum.Fail)
	fmt.Printf("Average latency: %v\n", sum.AvgLatency)
	fmt.Printf("Max latency: %v\n", sum.MaxLatency)
	fmt.Println("Status Codes:")
	for k, v := range sum.StatusCodes {
		fmt.Printf(" -%d: %d\n", k, v)
	}
	return nil
}

func help(ctx context.Context, args []string) error {
	for _, v := range cmds {
		fmt.Println("=============================================")
		fmt.Println(v.Name, ":")
		fmt.Println(v.Description)
		if v.Usage != "" {
			fmt.Println(v.Usage)
		}
		fmt.Print("=============================================\n\n")
	}
	return nil
}
