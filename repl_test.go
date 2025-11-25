package main

import "testing"

func TestCleanInput(t *testing.T) {

	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    " Hello WORld ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "Garchomp Charizard",
			expected: []string{"garchomp", "charizard"},
		},
		{
			input:    "Garchomp",
			expected: []string{"garchomp"},
		},
		{
			input:    " pikachu        bulbasaur CHARMANDER ",
			expected: []string{"pikachu", "bulbasaur", "charmander"},
		}, {
			input:    "pikachu",
			expected: []string{"pikachu"},
		}, {
			input:    "",
			expected: []string{""},
		},
		{
			input:    "    ",
			expected: []string{},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("Test fail: Slice mismatch. Expected %v got %v", c.expected, actual)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Test fail: Word mismatch. Expected %s got %s", expectedWord, word)
			}
		}
	}

}
