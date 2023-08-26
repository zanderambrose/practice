package utils

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:examplepassword@localhost:27017")
	c, err := mongo.Connect(CTX, clientOptions)
	if err != nil {
		return nil, err
	}

	err = c.Ping(CTX, nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	coll := client.Database("whoshittin").Collection(name)
	return coll
}
