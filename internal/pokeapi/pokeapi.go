package pokeapi

import (
	"encoding/json"
	"net/http"
)

type locationAreasResp struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous any    `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetLocationAreasResp(url string) (locationAreasResp, error) {
	res, err := http.Get(url)
	if err != nil {
		return locationAreasResp{}, err
	}
	defer res.Body.Close()

	var locAreas = locationAreasResp{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locAreas); err != nil {
		return locationAreasResp{}, err
	}

	var locationNames = []string{}
	for _, res := range locAreas.Results {
		locationNames = append(locationNames, res.Name)
	}

	return locAreas, nil

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
