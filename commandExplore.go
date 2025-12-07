package main

import (
	"fmt"

	"github.com/phpires/pokedexcli/internal/pokeapi"
)

func commandExplore(commandConfig *CommandConfig, userParams []string) error {
	if len(userParams) == 0 {
		return fmt.Errorf("must pass a region to explore command")
	}
	region := userParams[0]
	res := pokeapi.GetLocationAreaRegionV2(commandConfig.NextUrl, region)

	return nil
}
