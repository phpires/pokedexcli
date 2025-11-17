package main

import (
	"fmt"
	"strings"
)

func cleanInput(text string) []string {
	fmt.Printf("cleanInput(%s)\n", text)
	if len(text) == 0 {
		return []string{}
	}

	splittedText := strings.Fields(strings.ToLower(text))
	fmt.Printf("splittedText: %s\n", splittedText)

	return splittedText
}
