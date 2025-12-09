package main

import "fmt"

func commandPokedex(cfg *CommandConfig, userParams []string) error {

	if len(cfg.PokemonCaught) == 0 {
		fmt.Println("Pokedex empty. Go catch some.")
		return nil
	}
	fmt.Println("Your Pokedex:")
	for _, pokemonCaught := range cfg.PokemonCaught {
		fmt.Printf("  - %s\n", pokemonCaught.Name)
	}
	return nil
}
