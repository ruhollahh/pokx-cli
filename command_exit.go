package main

import (
	"fmt"
	"os"
)

func exitCommand(_ *config, args ...string) error {
	fmt.Println("Exiting the Pokx CLI...")
	os.Exit(0)
	return nil
}
