package main

import (
	"fmt"
	"strings"
)

func helpCommand(_ *config, args ...string) error {
	help := strings.Builder{}
	help.WriteString("Welcome to the Pokx CLI!\nUsage:\n\n")

	for _, cmd := range getCommands() {
		help.WriteString(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
	}

	fmt.Println(help.String())

	return nil
}
