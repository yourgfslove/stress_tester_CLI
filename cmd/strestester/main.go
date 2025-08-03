package main

import (
	"github.com/yourgfslove/stressTester/internal/application"
	"github.com/yourgfslove/stressTester/internal/commands"
	// "github.com/yourgfslove/stressTester/internal/config"
)

func main() {
	// cfg := config.MustLoadConfig()
	commands := commands.MustInitCommands()
	application.Start(commands)
	// TODO: inti cmds(maybe map or idk)
}
