package main

import (
	// "fmt"
	"strings"
)

func cleanInput(text string) []string {
	var cleanStrings []string
	trimmed := strings.TrimSpace(text)
	// fmt.Printf("trimmed: %#v\n", trimmed)

	lowerCased := strings.ToLower(trimmed)
	// fmt.Printf("lower cased: %#v\n", lowerCased)

	cleanStrings = strings.Fields(lowerCased)
	return cleanStrings
}
