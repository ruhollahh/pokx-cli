package main

import (
	"fmt"
	"strings"
)

func helpCommand(_ *config) error {
	help := strings.Builder{}
	_, err := help.WriteString("Welcome to the Pokx CLI!\nUsage:\n\n")
	if err != nil {
		return err
	}
	for _, cmd := range getCommands() {
		_, err = help.WriteString(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
		if err != nil {
			return err
		}
	}

	fmt.Println(help.String())

	return nil
}
