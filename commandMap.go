package main

import (
	"fmt"

	"github.com/phpires/pokedexcli/internal/pokeapi"
)

func commandMap(mapCommandConfig *MapCommandConfig) error {
	url := mapCommandConfig.NextUrl
	var response pokeapi.LocationAreaResponseJson
	var err error

	if url == "" {
		response, err = pokeapi.GetLocationAreaV2(0, 0)
	} else {
		response, err = pokeapi.GetLocationAreaV2(0, 0)
	}
	if err != nil {
		fmt.Printf("error requesting to poké api: %v", err)
	}

	mapCommandConfig.NextUrl = response.Next
	mapCommandConfig.PreviousUrl = response.Previous
	locationNames := extractLocationNameFromMapPokeApi(response.Results)

	for n := range locationNames {
		fmt.Println(locationNames[n])
	}

	return nil
}
