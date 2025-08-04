package commands

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/yourgfslove/stressTester/internal/aggregator"
	"github.com/yourgfslove/stressTester/internal/config"
	"github.com/yourgfslove/stressTester/internal/lib/validation"
	"github.com/yourgfslove/stressTester/internal/worker"
)

type Command struct {
	Name        string
	Description string
	Usage       string
	Callback    func(args []string) error
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
			Callback: func(args []string) error {
				return stressTest(MainConf, args)
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

func exit(args []string) error {
	os.Exit(0)
	return nil
}

func stressTest(cfg config.Config, args []string) error {
	fs := flag.NewFlagSet("stress", flag.ContinueOnError)

	var rps int
	var link string
	var method string
	var times int

	fs.IntVar(&rps, "rps", 10, "amount of request per second")
	fs.StringVar(&link, "link", "https://httpbin.org/get", "link on stress site")
	fs.StringVar(&method, "method", "GET", "Request method")
	fs.IntVar(&times, "times", 1, "num of seconds doing requests")

	if err := fs.Parse(args); err != nil {
		return fmt.Errorf("failed to parse parameters: %s", err.Error())
	}

	method = strings.ToUpper(method)

	if err := validation.StressParamsValidate(link, method); err != nil {
		return fmt.Errorf("bad params: %s", err.Error())
	}

	client := http.DefaultClient
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(times)*time.Second)
	defer cancel()
	jobs := make(chan *http.Request)
	results := make(chan worker.Result, 10)

	wg := &sync.WaitGroup{}
	for i := 0; i <= cfg.NumWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			worker.Run(ctx, client, results, jobs)
		}()
	}

	go func() {
		ticker := time.NewTicker(time.Second / time.Duration(rps))
		defer ticker.Stop()
		for {
			select {
			case <-ctx.Done():
				close(jobs)
				return
			case <-ticker.C:
				req, _ := http.NewRequest(method, link, nil)
				jobs <- req
			}
		}
	}()

	go func() {
		wg.Wait()
		close(results)
	}()
	summary := make(chan aggregator.StressSummary)
	go aggregator.Aggregator(summary, results)

	sum := <-summary
	fmt.Printf("Number of success requests: %d\n", sum.Success)
	fmt.Printf("Number of failed requests: %d\n", sum.Fail)
	fmt.Printf("Average latency: %v\n", sum.AvgLatency)
	fmt.Printf("Max latency: %v\n", sum.MaxLatency)
	fmt.Println("Status Codes:")
	for k, v := range sum.StatusCodes {
		fmt.Printf(" -%d: %d\n", k, v)
	}
	fmt.Println(time.Second / time.Duration(rps))
	return nil
}

func help(args []string) error {
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
