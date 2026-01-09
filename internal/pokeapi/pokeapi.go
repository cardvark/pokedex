package pokeapi

import (
	"encoding/json"
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

/*

Order:
- API call to get an decode the json response into a go struct
- commandExplore in repl.go to call the GetLocationAreaPokemonResp function and process the resp struct.
- update API call to cache response.


*/

func GetLocationAreaPokemonEncountersResp(url string) (locationAreaPokemonEncountersResp, error) {
	entryBytes, ok := pokecache.MemCache.Get(url)
	if ok {
		var locAreaPokeResp locationAreaPokemonEncountersResp
		err := json.Unmarshal(entryBytes, &locAreaPokeResp)
		return locAreaPokeResp, err
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreaPokemonEncountersResp{}, err
	}
	defer res.Body.Close()

	var locAreaPokeResp locationAreaPokemonEncountersResp
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locAreaPokeResp); err != nil {
		return locationAreaPokemonEncountersResp{}, err
	}

	resultBytes, err := json.Marshal(locAreaPokeResp)
	if err != nil {
		return locationAreaPokemonEncountersResp{}, err
	}

	pokecache.MemCache.Add(url, resultBytes)

	return locAreaPokeResp, nil

}

func GetLocationAreasResp(url string) (locationAreaResp, error) {
	entryBytes, ok := pokecache.MemCache.Get(url)
	if ok {
		// fmt.Printf("Retrieving from cache! Raw value: %v\n", entryBytes)
		var locAreasResp locationAreaResp
		err := json.Unmarshal(entryBytes, &locAreasResp)
		return locAreasResp, err
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreaResp{}, err
	}
	defer res.Body.Close()

	var locAreasResp = locationAreaResp{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locAreasResp); err != nil {
		return locationAreaResp{}, err
	}

	resultBytes, err := json.Marshal(locAreasResp)
	if err != nil {
		return locationAreaResp{}, err
	}

	pokecache.MemCache.Add(url, resultBytes)

	return locAreasResp, nil

}
