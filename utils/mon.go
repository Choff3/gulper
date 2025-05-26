package utils

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

func StoreBeers(beerString string) {

	beers := GetBeers(beerString)
	client, _ := mongo.Connect(options.Client().ApplyURI("mongodb://localhost:27017"))
	collection := client.Database("gulper").Collection("beers")

	for beer := range beers {
		_, err := collection.InsertOne(nil, beer)
		if err != nil {
			panic(err)
		}
	}
}
