package main

import "strings"

func cleanInput(input string) []string {
	lowerCased := strings.ToLower(input)
	return strings.Fields(lowerCased)
}
