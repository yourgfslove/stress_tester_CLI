package commands

import "os"

type command struct {
	name        string
	description string
	Callback    func(args []string) error
}

type Commands map[string]command

func MustInitCommands() Commands {
	commands := map[string]command{
		"exit": {
			name:        "exit",
			description: "exiting from cli",
			Callback:    exit,
		},
	}
	return commands
}

func exit(args []string) error {
	os.Exit(1)
	return nil
}
