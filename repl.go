package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokx > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]

		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		err := cmd.callback()
		if err != nil {
			fmt.Printf("something went wrong: %s\n", err)
			continue
		}
	}
}

func cleanInput(input string) []string {
	lowerCased := strings.ToLower(input)
	return strings.Fields(lowerCased)
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokx CLI",
			callback:    exitCommand,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    helpCommand,
		},
	}
}
