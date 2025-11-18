package main

import (
	"strings"
)

func CleanInput(text string) []string {
	if len(text) == 0 {
		return []string{""}
	}

	splittedText := strings.Fields(strings.ToLower(text))

	if splittedText == nil {
		return []string{""}
	}

	return splittedText
}
