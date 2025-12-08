package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(cfg *CommandConfig, userParams []string) error {

	if len(userParams) == 0 {
		return fmt.Errorf("must pass a pokemon name")
	}

	pkmName := userParams[0]

	pkm, err := cfg.PokeApiClient.GetPokemon(pkmName)

	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", pkmName)

	ok, err := tryCatch(pkm.BaseExperience)

	if err != nil {
		return err
	}

	if ok {
		fmt.Printf("%v was caught!\n", pkmName)
		cfg.PokemonCaught[pkm.Name] = pkm
		return nil
	}

	fmt.Printf("%v escaped!\n", pkmName)
	return nil
}

func tryCatch(baseExp int) (bool, error) {

	catchChance := float64(rand.Intn(baseExp)) / float64(baseExp)
	if catchChance < 0.7 {
		return false, nil
	}
	return true, nil
}
