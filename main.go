package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		var input []string
		if scanner.Scan() {
			input = CleanInput(scanner.Text())
			fmt.Printf("Your command was: %s\n", input[0])
		}
	}
}
