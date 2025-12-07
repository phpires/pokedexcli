package main

import (
	"time"

	"github.com/phpires/pokedexcli/internal/pokeapi"
)

func main() {

	commandConfig := CommandConfig{
		PokeApiClient: pokeapi.NewClient(5*time.Second, 5*time.Minute),
		NextUrl:       "",
		PreviousUrl:   "",
	}
	startRepl(&commandConfig)
}
