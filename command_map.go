package main

import (
	"fmt"

	"github.com/phpires/pokedexcli/internal/pokeapi"
)

func commandMap(commandConfig *CommandConfig, userParams []string) error {
	fmt.Println("Executing map command.")
	response, err := commandConfig.PokeApiClient.GetLocations(commandConfig.NextUrl)

	if err != nil {
		fmt.Printf("error requesting to poké api: %v", err)
	}

	updateUrlsMapCommandConfig(commandConfig, response)
	printAnswer(response.Results)

	return nil
}

func commandMapB(commandConfig *CommandConfig, userParams []string) error {

	response, err := commandConfig.PokeApiClient.GetLocations(commandConfig.PreviousUrl)

	if err != nil {
		fmt.Printf("error requesting to poké api: %v", err)
	}

	updateUrlsMapCommandConfig(commandConfig, response)
	printAnswer(response.Results)

	return nil
}

func updateUrlsMapCommandConfig(mapCommandConfig *CommandConfig, response pokeapi.Location) {
	mapCommandConfig.NextUrl = response.Next
	mapCommandConfig.PreviousUrl = response.Previous
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
