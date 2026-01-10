package main

import (
	"bufio"
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/cardvark/pokedex/internal/pokeapi"
)

type CliCommand struct {
	name        string
	description string
	callback    func(*Config, string) error
}

type Config struct {
	locAreasNext  string
	locAreasPrev  string
	caughtPokemon map[string]pokeapi.Pokemon
}

const locationAreaURL string = "https://pokeapi.co/api/v2/location-area/"
const pokemoneURL string = "https://pokeapi.co/api/v2/pokemon/"

func startRepl() {
	scanner := bufio.NewScanner(os.Stdin)
	mapConfig := Config{
		locAreasNext:  locationAreaURL,
		locAreasPrev:  "",
		caughtPokemon: map[string]pokeapi.Pokemon{},
	}

	for {
		fmt.Printf("Pokedex > ")

		if scanner.Scan() {

			line := scanner.Text()

			cleanSlice := cleanInput(line)
			if len(cleanSlice) < 1 {
				continue
			}

			command := cleanSlice[0]
			paramOne := ""
			if len(cleanSlice) > 1 {
				paramOne = cleanSlice[1]
			}

			if value, ok := getCommands()[command]; ok {
				err := value.callback(&mapConfig, paramOne)
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
		"explore": {
			name:        "explore",
			description: "Display list of pokemon in a given area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Attempt to catch a pokemon by name",
			callback:    commandCatch,
		},
	}
}

func commandCatch(cfg *Config, pokemonName string) error {
	if pokemonName == "" {
		return errors.New("No pokemon name provided.")
	}

	fullUrl := pokemoneURL + pokemonName
	pokemon, err := pokeapi.GetPokemon(fullUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name)

	baseChance := 45.0
	randInt := rand.Intn(pokemon.BaseExperience)
	result := baseChance / float64(randInt)

	if result >= 1 {
		fmt.Printf("%s was caught!\n", pokemon.Name)
		cfg.caughtPokemon[pokemon.Name] = pokemon
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
	}

	return nil
}

func commandExplore(_ *Config, areaName string) error {
	fullUrl := locationAreaURL + areaName
	fmt.Printf("Exploring %s...\n", areaName)

	pokeEncResp, err := pokeapi.GetLocationAreaPokemonEncountersResp(fullUrl)
	if err != nil {
		return err
	}

	fmt.Printf("Found Pokemon:\n")

	for _, pokeStruct := range pokeEncResp.PokemonEncounters {
		fmt.Printf("- %s\n", pokeStruct.Pokemon.Name)
	}

	return nil
}

func commandExit(_ *Config, _ string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(_ *Config, _ string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Printf("Usage:\n\n")

	for key, cmdStruct := range getCommands() {
		fmt.Printf("%s: %s\n", key, cmdStruct.description)
	}
	return nil
}

func commandMap(cfg *Config, _ string) error {
	if cfg.locAreasNext == "" {
		return fmt.Errorf("No subsequent area locations.")
	}

	locAreas, err := pokeapi.GetLocationAreasResp(cfg.locAreasNext)

	if err != nil {
		return err
	}

	cfg.locAreasPrev = cfg.locAreasNext
	cfg.locAreasNext = locAreas.Next

	for _, loc := range locAreas.Results {
		fmt.Println(loc.Name)
	}

	return nil
}

func commandMapb(cfg *Config, _ string) error {
	if cfg.locAreasPrev == "" {
		fmt.Println("You're on the first page.")
		return nil
	}

	locAreas, err := pokeapi.GetLocationAreasResp(cfg.locAreasPrev)

	if err != nil {
		return err
	}

	prevString := fmt.Sprint(locAreas.Previous)
	// fmt.Println(prevString)

	if prevString == "<nil>" {
		cfg.locAreasPrev = ""
	} else {
		cfg.locAreasPrev = prevString
	}

	cfg.locAreasNext = locAreas.Next

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
