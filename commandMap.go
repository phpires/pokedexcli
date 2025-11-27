package main

import (
	"fmt"

	"github.com/phpires/pokedexcli/internal/pokeapi"
)

func commandMap(mapCommandConfig *MapCommandConfig) error {

	response, err := pokeapi.GetLocationAreaV2(mapCommandConfig.NextUrl)
	if err != nil {
		fmt.Printf("error requesting to poké api: %v", err)
	}

	updateUrlsMapCommandConfig(mapCommandConfig, response)
	printAnswer(response.Results)
	return nil
}

func updateUrlsMapCommandConfig(mapCommandConfig *MapCommandConfig, response pokeapi.LocationAreaResponseJson) {
	mapCommandConfig.NextUrl = response.Next
	mapCommandConfig.PreviousUrl = response.Previous
}

func commandMapB(mapCommandConfig *MapCommandConfig) error {
	response, err := pokeapi.GetLocationAreaV2(mapCommandConfig.PreviousUrl)
	if err != nil {
		fmt.Printf("error requesting to poké api: %v", err)
	}

	updateUrlsMapCommandConfig(mapCommandConfig, response)
	printAnswer(response.Results)
	return nil
}

func printAnswer(results []pokeapi.Results) {
	locationNames := extractLocationsResponse(results)

	for n := range locationNames {
		fmt.Println(locationNames[n])
	}
}

func extractLocationsResponse(results []pokeapi.Results) []string {
	response := make([]string, len(results))
	for i := range results {
		response[i] = results[i].Name
	}
	return response
}
