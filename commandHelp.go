package main

import (
	"fmt"
	"strings"
)

func commandHelp(mapCommandConfig *CommandConfig, userParams []string) error {
	var builder strings.Builder
	builder.WriteString("Welcome to the Pokedex!\n")
	builder.WriteString("Usage:\n")
	builder.WriteString("\n")
	commands := getCommands()
	for _, cmd := range commands {
		builder.WriteString(fmt.Sprintf("%s: %s\n", cmd.name, cmd.description))
	}
	fmt.Println(builder.String())
	return nil
}
