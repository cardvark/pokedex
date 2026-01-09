package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/cardvark/pokedex/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config) error
}

type Config struct {
	next     string
	previous string
}

const locationAreasURL string = "https://pokeapi.co/api/v2/location-area/"

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	mapConfig := Config{next: locationAreasURL, previous: ""}

	for {
		fmt.Printf("Pokedex > ")

		if scanner.Scan() {

			line := scanner.Text()

			cleanSlice := cleanInput(line)
			command := cleanSlice[0]

			if value, ok := getCommands()[command]; ok {
				err := value.callback(&mapConfig)
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

func getCommands() map[string]CliCommand {
	return map[string]CliCommand{
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
		"map": {
			name:        "map",
			description: "Provides a list of the next 20 areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns to previous 20 areas.",
			callback:    commandMapb,
		},
	}
}

func commandExit(cfg *Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for key, cmdStruct := range getCommands() {
		fmt.Printf("%s: %s\n", key, cmdStruct.description)
	}
	return nil
}

func commandMap(cfg *Config) error {
	if cfg.next == "" {
		return fmt.Errorf("No subsequent area locations.")
	}

	locAreas, err := pokeapi.GetLocationAreaList(cfg.next)

	if err != nil {
		return err
	}

	cfg.previous = cfg.next
	cfg.next = locAreas.Next

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *Config) error {
	if cfg.previous == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	locAreas, err := pokeapi.GetLocationAreaList(cfg.previous)

	if err != nil {
		return err
	}

	prevString := fmt.Sprint(locAreas.Previous)
	// fmt.Println(prevString)

	if prevString == "<nil>" {
		cfg.previous = ""
	} else {
		cfg.previous = prevString
	}

	cfg.next = locAreas.Next

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
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
