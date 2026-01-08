package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("Pokedex > ")

		if scanner.Scan() {

			line := scanner.Text()

			cleanSlice := cleanInput(line)
			command := cleanSlice[0]

			if value, ok := getCommands()[command]; ok {
				err := value.callback()
				if err != nil {
					fmt.Println(err)
				}
				continue
			} else {
				fmt.Println("Unknown command.")
				continue
			}
		}
	}
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
	}
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for key, cmdStruct := range commandMap {
		fmt.Printf("%s: %s\n", key, cmdStruct.description)
	}
	return nil
}

func cleanInput(text string) []string {
	var cleanStrings []string
	trimmed := strings.TrimSpace(text)
	lowerCased := strings.ToLower(trimmed)
	cleanStrings = strings.Fields(lowerCased)
	return cleanStrings
}
