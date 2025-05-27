package utils

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Beer struct {
	Name    string  `json:"name" bson:"name"`
	Brewery string  `json:"brewery" bson:"brewery"`
	Style   string  `json:"style" bson:"style"`
	Abv     float64 `json:"abv" bson:"abv"`
	Price   float64 `json:"price" bson:"price"`
	Venue   string  `json:"venue" bson:"venue"`
	Url     string  `json:"url" bson:"url"`
}

func Gulp(beerStr, venue, website string, store bool) []Beer {

	beers := convertString(beerStr)

	if store {
		fmt.Printf("Storing %d beers for %s\n", len(beers), venue)
		storeBeers(beers, venue, website)
	}
	return beers
}

func convertString(beerString string) []Beer {

	var beers []Beer

	jsonBytes := []byte(beerString)

	err := json.Unmarshal(jsonBytes, &beers)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON array: %v", err)
	}

	return beers
}

func storeBeers(beers []Beer, venue, url string) {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("Error: MONGODB_URI environment variable not set. Please set it before running.")
	}
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	coll := client.Database("gulper").Collection("beers")

	for _, beer := range beers {
		beer.Venue = venue
		beer.Url = url
		_, err := coll.InsertOne(context.TODO(), beer)
		if err != nil {
			panic(err)
		}
	}
}

func GetUserAgent() string {
	return "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:138.0) Gecko/20100101 Firefox/138.0"
}
