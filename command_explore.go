package main

import "fmt"

func exploreCommand(cfg *config, args ...string) error {
	if len(args) != 1 {
		return fmt.Errorf("area can't be empty")
	}

	area := args[0]
	fmt.Printf("Exploring %s...\n", area)
	pokemons, err := cfg.pokeapiClient.GetPokemons(area)
	if err != nil {
		return fmt.Errorf("listPokemon: %w", err)
	}

	fmt.Println("Found Pokemons:")
	for _, e := range pokemons.PokemonEncounters {
		fmt.Println(e.Pokemon.Name)
	}

	return nil
}
