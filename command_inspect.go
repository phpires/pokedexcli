package main

import "fmt"

func commandInspect(cfg *CommandConfig, userParams []string) error {
	if len(userParams) == 0 {
		return fmt.Errorf("must pass a region to explore command")
	}

	pokemon_name := userParams[0]
	pokemon, ok := cfg.PokemonCaught[pokemon_name]
	if !ok {
		fmt.Println("you have not caught that pokemon")
		return nil
	}

	fmt.Printf("Name: %v\n", pokemon.Name)
	fmt.Printf("Height: %v\n", pokemon.Height)
	fmt.Printf("Weight: %v\n", pokemon.Weight)

	fmt.Println("Stats:")
	for _, s := range pokemon.Stats {
		fmt.Printf("  -%s: %d\n", s.Stat.Name, s.BaseStat)

	}

	fmt.Println("Types:")
	for _, t := range pokemon.Types {
		fmt.Printf("  - %s\n", t.Type.Name)

	}

	return nil
}
