package pokeapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/cardvark/pokedex/internal/pokecache"
)

type locationAreaResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

type locationAreaPokemonEncountersResp struct {
	ID       int `json:"id"`
	Location struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name              string `json:"name"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type Pokemon struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Stats          []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"stat"`
	} `json:"stats"`
	Types []struct {
		Slot int `json:"slot"`
		Type struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"type"`
	} `json:"types"`
	Height int `json:"height"`
	Weight int `json:"weight"`
}

func GetResource[T any](url string) (T, error) {
	var result T

	entryBytes, ok := pokecache.MemCache.Get(url)
	if ok {
		err := json.Unmarshal(entryBytes, &result)
		return result, err
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return result, err
	}

	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	if string(dat) == "Not Found" {
		return result, errors.New("Resource not found. Please check your spelling and try again.")
	}

	pokecache.MemCache.Add(url, dat)

	err = json.Unmarshal(dat, &result)
	if err != nil {
		fmt.Printf("Error unmarshaling data: %v", err)
	}
	return result, err
}

// func GetResource[T any](url string) (T, error) {
// 	var result T

// 	entryBytes, ok := pokecache.MemCache.Get(url)
// 	if ok {
// 		err := json.Unmarshal(entryBytes, &result)
// 		return result, err
// 	}

// 	res, err := http.Get(url)
// 	if err != nil {
// 		return result, err
// 	}
// 	defer res.Body.Close()

// 	dat, err := io.ReadAll(res.Body)
// 	if err != nil {
// 		return result, err
// 	}

// 	pokecache.MemCache.Add(url, dat)

// 	err = json.Unmarshal(dat, &result)
// 	return result, err
// }

func GetPokemon(url string) (Pokemon, error) {
	return GetResource[Pokemon](url)
}

func GetLocationAreaPokemonEncountersResp(url string) (locationAreaPokemonEncountersResp, error) {
	return GetResource[locationAreaPokemonEncountersResp](url)
}

func GetLocationAreasResp(url string) (locationAreaResp, error) {
	return GetResource[locationAreaResp](url)
}
