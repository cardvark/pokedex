package pokeapi

import (
	"encoding/json"
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

type pokemonResp struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
}

func GetResource[T any](url string) (T, error) {
	var result T

	entryBytes, ok := pokecache.MemCache.Get(url)
	if ok {
		err := json.Unmarshal(entryBytes, &result)
		return result, err
	}

	res, err := http.Get(url)
	if err != nil {
		return result, err
	}
	defer res.Body.Close()

	dat, err := io.ReadAll(res.Body)
	if err != nil {
		return result, err
	}

	pokecache.MemCache.Add(url, dat)

	err = json.Unmarshal(dat, &result)
	return result, err
}

func GetPokemonResp(url string) (pokemonResp, error) {
	return GetResource[pokemonResp](url)
}

func GetLocationAreaPokemonEncountersResp(url string) (locationAreaPokemonEncountersResp, error) {
	return GetResource[locationAreaPokemonEncountersResp](url)
}

func GetLocationAreasResp(url string) (locationAreaResp, error) {
	return GetResource[locationAreaResp](url)
}
