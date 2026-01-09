package pokeapi

import (
	"fmt"
	"reflect"
	"testing"
)

func TestGetLocationAreaList(t *testing.T) {
	cases := map[string]struct {
		input         string
		expectedSlice []string
		expectedNext  string
		expectedErr   error
	}{
		"base test": {
			input:         "https://pokeapi.co/api/v2/location-area/",
			expectedSlice: []string{"canalave-city-area", "eterna-city-area", "pastoria-city-area", "sunyshore-city-area", "sinnoh-pokemon-league-area", "oreburgh-mine-1f", "oreburgh-mine-b1f", "valley-windworks-area", "eterna-forest-area", "fuego-ironworks-area", "mt-coronet-1f-route-207", "mt-coronet-2f", "mt-coronet-3f", "mt-coronet-exterior-snowfall", "mt-coronet-exterior-blizzard", "mt-coronet-4f", "mt-coronet-4f-small-room", "mt-coronet-5f", "mt-coronet-6f", "mt-coronet-1f-from-exterior"},
			expectedNext:  "https://pokeapi.co/api/v2/location-area/?offset=20&limit=20",
			expectedErr:   nil,
		},
	}

	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			locAreas, actualErr := GetLocationAreaList(tc.input)
			fmt.Printf("Actual slice: %v\n", locAreas.Results)
			fmt.Printf("Actual next: %v\n", locAreas.Next)

			actualSlice := []string{}

			for _, loc := range locAreas.Results {
				actualSlice = append(actualSlice, loc.Name)
			}

			if !reflect.DeepEqual(tc.expectedSlice, actualSlice) {
				t.Fatalf("expected: %#v, actual: %#v", tc.expectedSlice, actualSlice)
			}
			if !reflect.DeepEqual(tc.expectedNext, locAreas.Next) {
				t.Fatalf("expected: %#v, actual: %#v", tc.expectedNext, locAreas.Next)
			}
			if !reflect.DeepEqual(tc.expectedErr, actualErr) {
				t.Fatalf("expected: %#v, actual: %#v", tc.expectedErr, actualErr)
			}
		})
	}

}
