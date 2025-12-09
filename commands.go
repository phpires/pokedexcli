package main

import "github.com/phpires/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*CommandConfig, []string) error
}

type CommandConfig struct {
	PokeApiClient pokeapi.Client
	NextUrl       string
	PreviousUrl   string
	RegionName    string
	PokemonCaught map[string]pokeapi.PokemonInfo
}

func getCommands() map[string]cliCommand {
	commands := map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Display the next maps locations",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Display the previous maps locations",
			callback:    commandMapB,
		},
		"explore": {
			name:        "explore <region_name>",
			description: "Explore a specified location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Catch a pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "inspect a caught pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Lst pokedex entries",
			callback:    commandPokedex,
		},
	}
	return commands
}
