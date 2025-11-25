package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type MapCommandConfig struct {
	NextUrl     string
	PreviousUrl string
}

func startRepl() {

	scanner := bufio.NewScanner(os.Stdin)

	mapCommandConfig := MapCommandConfig{
		NextUrl:     "",
		PreviousUrl: "",
	}

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		userCommand := input[0]
		fmt.Printf("UserCommand: %v\n", userCommand)
		cmd, commandFound := getCommands()[userCommand]
		if commandFound {
			err := cmd.callback(&mapCommandConfig)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}

}

func cleanInput(text string) []string {
	if len(text) == 0 {
		return []string{""}
	}

	splittedText := strings.Fields(strings.ToLower(text))

	if splittedText == nil {
		return []string{""}
	}

	return splittedText
}
