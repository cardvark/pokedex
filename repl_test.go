package main

import (
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := map[string]struct {
		input    string
		expected []string
	}{
		"extra spaces": {
			input:    "  hello  world   ",
			expected: []string{"hello", "world"},
		},
		"capitalization": {
			input:    "Charmander Bulbasaur PIKACHU",
			expected: []string{"charmander", "bulbasaur", "pikachu"},
		},
		"one word split": {
			input:    " hello    ",
			expected: []string{"hello"},
		},
		"just whitespace": {
			input:    "   ",
			expected: []string{},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("expected: %#v, actual: %#v", tc.expected, actual)
			}
		})
	}

}

// TODO: test for command inspect.
// use caterpie; lower than current baseChance 45.
