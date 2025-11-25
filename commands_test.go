package main

import (
	"testing"
)

func TestGetCommands(t *testing.T) {

	cases := []struct {
		commandsSize  int
		commandsNames []string
	}{
		{
			commandsSize:  4,
			commandsNames: []string{"exit", "help", "map", "mapb"},
		},
	}

	for _, c := range cases {
		actual := getCommands()

		if c.commandsSize != len(actual) {
			t.Errorf("Test fail: Commands size is wrong. Expected %v got %v", c.commandsSize, len(actual))
		}

		for idx := range c.commandsNames {
			cmd, ok := actual[c.commandsNames[idx]]
			if !ok {
				t.Errorf("Test fail: Command '%v' not found", c.commandsNames[idx])
			}
			if cmd.name != c.commandsNames[idx] {
				t.Errorf("Test fail: Command '%v' with wrong name.\n Should be '%v' but found '%v'", c.commandsNames[idx], c.commandsNames[idx], cmd.name)
			}
		}

	}
}
