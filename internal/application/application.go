package application

import (
	"bufio"
	"fmt"
	"os"

	"github.com/yourgfslove/stressTester/internal/commands"
	"github.com/yourgfslove/stressTester/internal/lib/input"
)

func Start(commands commands.Commands) error {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Stress >")
		if scanner.Scan() {
			input := input.Clean(scanner.Text())
			if len(input) == 0 {
				continue
			}
			if command, ok := commands[input[0]]; ok {
				err := command.Callback(input[1:])
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}
