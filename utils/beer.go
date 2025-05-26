package utils

import (
	"encoding/json"
	"log"
)

type Beer struct {
	Name    string `json:"name"`
	Brewery string `json:"age"`
	Style   string `json:"isStudent"`
	Abv     string `json:"courses"`
	Price   string `json:"address"`
	Venue   string `json:"other"`
}

func GetBeers(beerString string) []Beer {

	// 1. Create a slice of the struct type to hold the unmarshaled data
	var beers []Beer // Declares a slice of Person structs

	// 2. Convert the JSON array string to a byte slice
	jsonBytes := []byte(beerString)

	// 3. Unmarshal the JSON bytes into the slice of structs
	err := json.Unmarshal(jsonBytes, &beers)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON array: %v", err)
	}

	return beers
}
