package main

import (
	"fmt"
	"math/rand"
)

func catchCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("name can't be empty")
	}
	name := args[0]

	if _, exists := cfg.caughtPokemons[name]; exists {
		fmt.Printf("%s is already caught\n", name)
		return nil
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", name)

	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return fmt.Errorf("getPokemon: %w", err)
	}

	randomValue := rand.Intn(pokemon.BaseExperience)

	if randomValue > 40 {
		fmt.Printf("%s escaped!\n", name)
		return nil
	}

	fmt.Printf("%s was caught!\n", name)
	cfg.caughtPokemons[name] = pokemon

	return nil
}
