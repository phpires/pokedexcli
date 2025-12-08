package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(commandConfig *CommandConfig) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := cleanInput(scanner.Text())
		userCommand := input[0]
		userParams := input[1:]

		cmd, commandFound := getCommands()[userCommand]
		if commandFound {
			err := cmd.callback(commandConfig, userParams)
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
