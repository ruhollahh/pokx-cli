package main

import (
	"fmt"
	"strings"
)

func inspectCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("name can't be empty")
	}
	name := args[0]

	pokemon, exists := cfg.caughtPokemons[name]
	if !exists {
		fmt.Printf("You haven't caught %s yet!\n", name)
		return nil
	}

	pokemonInfo := strings.Builder{}
	pokemonInfo.WriteString(fmt.Sprintf("Name: %s\nHeight: %d\nWeight: %d\n", pokemon.Name, pokemon.Height, pokemon.Weight))

	pokemonInfo.WriteString("Stats:\n")
	for _, s := range pokemon.Stats {
		pokemonInfo.WriteString(fmt.Sprintf("  -%s: %d\n", s.Stat.Name, s.BaseStat))
	}

	pokemonInfo.WriteString("Types:\n")
	for _, t := range pokemon.Types {
		pokemonInfo.WriteString(fmt.Sprintf("  - %s\n", t.Type.Name))
	}

	fmt.Println(pokemonInfo.String())

	return nil
}
