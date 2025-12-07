package main

import "github.com/phpires/pokedexcli/internal/pokeapi"

type cliCommand struct {
	name        string
	description string
	callback    func(*CommandConfig) error
}

type CommandConfig struct {
	PokeApiClient pokeapi.Client
	NextUrl       string
	PreviousUrl   string
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
	}
	return commands
}
