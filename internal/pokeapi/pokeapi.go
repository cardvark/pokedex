package pokeapi

import (
	"encoding/json"
	"net/http"

	"github.com/cardvark/pokedex/internal/pokecache"
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
	// TODO: check if entry exists in cacheMap;
	// if no, then move on.
	// if yes, retrieve the entry from the cacheMap
	// convert the cacheEntry value from bytes into locationAreasResp struct.
	// declare new locationsAreasResp struct
	// create
	// decode
	// return locAreasResp

	entryBytes, ok := pokecache.MemCache.Get(url)
	if ok {
		// fmt.Printf("Retrieving from cache! Raw value: %v\n", entryBytes)
		var locAreasResp locationAreasResp
		err := json.Unmarshal(entryBytes, &locAreasResp)
		return locAreasResp, err
	}

	res, err := http.Get(url)
	if err != nil {
		return locationAreasResp{}, err
	}
	defer res.Body.Close()

	var locAreasResp = locationAreasResp{}
	decoder := json.NewDecoder(res.Body)

	if err := decoder.Decode(&locAreasResp); err != nil {
		return locationAreasResp{}, err
	}

	// TODO: put new value into cachemap.
	// encode using gob to bytes
	// create a buffer to hold gob-encoded bytes
	// create an encoder and send the data
	// extract the bytes

	resultBytes, err := json.Marshal(locAreasResp)
	if err != nil {
		return locationAreasResp{}, err
	}

	pokecache.MemCache.Add(url, resultBytes)

	return locAreasResp, nil

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
