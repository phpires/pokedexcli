package main

import (
	"fmt"
)

func commandExplore(commandConfig *CommandConfig, userParams []string) error {
	if len(userParams) == 0 {
		return fmt.Errorf("must pass a region to explore command")
	}

	commandConfig.RegionName = userParams[0]

	res, err := commandConfig.PokeApiClient.GetLocationAreaRegionV2(commandConfig.RegionName)
	if err != nil {
		return err
	}

	for _, pkmEncounter := range res.PokemonEncounters {
		fmt.Println(pkmEncounter.Pokemon.Name)
	}

	return nil
}
