package commands

import (
	"fmt"
	"os"

	"github.com/yourgfslove/stressTester/internal/config"
)

type Command struct {
	name        string
	description string
	Usage       string
	Callback    func(args []string) error
}

type Commands map[string]Command

var cfg config.Config
var cmds map[string]Command

func MustInitCommands(MainConf config.Config) Commands {
	cfg = MainConf
	cmds = map[string]Command{
		"exit": {
			name:        "exit",
			description: "exiting from cli",
			Callback:    exit,
		},
		"stress": {
			name:        "Stress",
			description: "starting stress test",
			Usage:       " --rps - amount of request per second \n --link - link on stress site \n --method[GET/POST/PATCH] - method",
			Callback:    stressTest,
		},
		"help": {
			name:        "Help",
			description: "describing all cmds",
			Callback:    help,
		},
	}
	return cmds
}

func exit(args []string) error {
	os.Exit(0)
	return nil
}

func stressTest(args []string) error {
	panic("imp me")
}

func help(args []string) error {
	for key, v := range cmds {
		fmt.Println("=============================================")
		fmt.Println(key, ":")
		fmt.Println(v.description)
		if v.Usage != "" {
			fmt.Println(v.Usage)
		}
		fmt.Print("=============================================\n\n")
	}
	return nil
}
