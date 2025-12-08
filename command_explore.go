package main

import (
	"fmt"
)

func commandExplore(commandConfig *CommandConfig, userParams []string) error {
	if len(userParams) == 0 {
		return fmt.Errorf("must pass a region to explore command")
	}

	commandConfig.RegionName = userParams[0]

	location, err := commandConfig.PokeApiClient.GetLocationRegion(commandConfig.RegionName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %v...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, pkmEncounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", pkmEncounter.Pokemon.Name)
	}

	return nil
}
