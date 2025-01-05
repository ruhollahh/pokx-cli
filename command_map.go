package main

import (
	"fmt"
)

func mapCommand(cfg *config, args ...string) error {
	areas, err := cfg.pokeapiClient.ListAreas(cfg.next)
	if err != nil {
		return fmt.Errorf("listAreas: %w", err)
	}

	cfg.next = areas.Next
	cfg.previous = areas.Previous

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}

func mapBackCommand(cfg *config, args ...string) error {
	if cfg.previous == nil {
		fmt.Println("you're on the first page")
		return nil
	}

	areas, err := cfg.pokeapiClient.ListAreas(cfg.previous)
	if err != nil {
		return fmt.Errorf("listAreas: %w", err)
	}

	cfg.next = areas.Next
	cfg.previous = areas.Previous

	for _, a := range areas.Results {
		fmt.Println(a.Name)
	}

	return nil
}
