package utils

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func InitDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:examplepassword@db:27017")
	client, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		return nil, err
	}

	mongoClient = client

	err = mongoClient.Ping(CTX, nil)
	if err != nil {
		return nil, err
	}

	return mongoClient, nil
}

func GetMongoClient() (*mongo.Client, error) {
	if mongoClient == nil {
		fmt.Println("HAVING TO INIT DB")
		return InitDB()
	}

	fmt.Println("RETURNING SINGLE MONGO CLIENT")
	return mongoClient, nil
}

func GetCollection(name string) *mongo.Collection {
	client, err := GetMongoClient()
	if err != nil {
		fmt.Println("Error getting mongo client in get collection")
	}
	coll := client.Database("whoshittin").Collection(name)
	return coll
}
