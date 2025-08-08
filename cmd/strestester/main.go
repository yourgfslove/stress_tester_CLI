package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/yourgfslove/stress_tester_CLI/internal/application"
	"github.com/yourgfslove/stress_tester_CLI/internal/commands"
	"github.com/yourgfslove/stress_tester_CLI/internal/config"
)

func main() {
	cfg := config.MustLoadConfig()
	commands := commands.MustInitCommands(*cfg)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-stop
		cancel()
		fmt.Println("interupted")
	}()
	if err := application.Start(ctx, commands); err != nil {
		fmt.Println("Error", err)
	}
}
