package main

import (
	"bufio"
	"fmt"
	"github.com/ruhollahh/pokx-cli/internal/clients/pokeapi"
	"os"
	"strings"
)

type config struct {
	pokeapiClient  *pokeapi.Client
	caughtPokemons map[string]pokeapi.RespPokemon
	next           *string
	previous       *string
}

func startRepl(cfg *config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokx > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		if len(input) == 0 {
			continue
		}
		commandName := input[0]
		var args []string
		if len(input) > 1 {
			args = input[1:]
		}

		cmd, exists := getCommands()[commandName]
		if !exists {
			fmt.Printf("Unknown command: %s\n", commandName)
			continue
		}

		err := cmd.callback(cfg, args...)
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
	callback    func(*config, ...string) error
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
		"map": {
			name:        "map",
			description: "Displays a paginated list of all the location areas, call again to display the next page",
			callback:    mapCommand,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the previous page of the location areas",
			callback:    mapBackCommand,
		},
		"explore": {
			name:        "explore <area>",
			description: "Displays the pokemons in an area",
			callback:    exploreCommand,
		},
		"catch": {
			name:        "catch <pokemon name>",
			description: "Try to catch a pokemon",
			callback:    catchCommand,
		},
		"inspect": {
			name:        "inspect <pokemon name>",
			description: "Inspect a pokemon that you've already caught",
			callback:    inspectCommand,
		},
	}
}
