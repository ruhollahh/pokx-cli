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

		fmt.Printf("Your command was: %s\n", input[0])
	}
}

func cleanInput(input string) []string {
	lowerCased := strings.ToLower(input)
	return strings.Fields(lowerCased)
}
