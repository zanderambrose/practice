package utils

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI("mongodb://root:examplepassword@localhost:27017")
	c, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = c.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func InsertOne(person *Person, coll *mongo.Collection) (*mongo.InsertOneResult, error) {
	result, err := coll.InsertOne(CTX, person)
	if err != nil {
		fmt.Println(err)
	}
	return result, err
}

func GetCollection(client *mongo.Client, name string) *mongo.Collection {
	coll := client.Database("practice").Collection(name)
	return coll
}
