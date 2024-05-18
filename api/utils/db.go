package db

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

var CTX = context.Background()

func InitDB() (*mongo.Client, error) {
	// TODO - This needs env variable
	clientOptions := options.Client().ApplyURI("mongodb://root:examplepassword@db:27017")
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		// TODO - Log handling
		return nil, err
	}

	mongoClient = client

	err = mongoClient.Ping(CTX, nil)
	if err != nil {
		// TODO - Log handling
		return nil, err
	}

	return mongoClient, nil
}

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		return InitDB()
	}
	return mongoClient, nil
}

func GetDatabase() *mongo.Database {
	client, err := GetMongoClient()
	// ERROR HANDLE
	if err != nil {
		fmt.Println("Error getting mongo client in get collection")
	}

	// TODO - This needs env variable
	return client.Database("whoshittin")
}

func GetCollection(name string) *mongo.Collection {
	db := GetDatabase()
	return db.Collection(name)
}

func FilterCollectionResults(collections []string) []string {
	var collectionsToExclude = []string{"client"}
	var filteredCollections []string
	for _, item := range collections {
		shouldInclude := true
		for _, itemToExclude := range collectionsToExclude {
			if item == itemToExclude {
				shouldInclude = false
			}
		}
		if shouldInclude {
			filteredCollections = append(filteredCollections, item)
		}
	}
	return filteredCollections
}
