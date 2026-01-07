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
			expected: []string{"charmander, bulbasaur, pikachu"},
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			actual := cleanInput(tc.input)
			if !reflect.DeepEqual(tc.expected, actual) {
				t.Fatalf("expected: %v, actual: %v", tc.expected, actual)
			}
		})
	}

}
