package pokeapi

import (
	"encoding/json"
	"net/http"
)

const locationAreaUrl string = "https://pokeapi.co/api/v2/location-area/"

type locationAreas struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func getLocationAreaList(url string) ([]string, string, error) {
	res, err := http.Get(url)
	if err != nil {
		return []string{}, "", err
	}
	defer res.Body.Close()

	var locAreas = locationAreas{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locAreas); err != nil {
		return []string{}, "", err
	}

	var locationNames = []string{}
	for _, res := range locAreas.Results {
		locationNames = append(locationNames, res.Name)
	}

	return locationNames, locAreas.Next, nil

}

/*
Walk through:
enter "map" in the terminal.
	"map" command added to commandMap function.
	"map" command function created:
		Calls ...

Get list of areas from poke api, offset.
first call, offset 0. subsequent calls offset by 20.
	where am I tracking the offset?


*/
