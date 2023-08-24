package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"practice/pkg/models"
	"practice/pkg/utils"
)

type MongoClient struct {
	MongoClient *mongo.Client
}

func InitClient(client *mongo.Client) *MongoClient {
	return &MongoClient{
		MongoClient: client,
	}
}

func (client *MongoClient) GetAllUser(c *fiber.Ctx) error {
	cursor, err := utils.GetCollection(client.MongoClient, "people").Find(utils.CTX, bson.D{})
	if err != nil {
		fmt.Println("Error finding documents")
	}

	defer cursor.Close(utils.CTX)

	var results []models.Person
	if err := cursor.All(utils.CTX, &results); err != nil {
		fmt.Println("Error ", err)
	}

	if err := cursor.Err(); err != nil {
		fmt.Println("Error ", err)
	}

	response := map[string]interface{}{
		"people": results,
	}
	return c.JSON(response)
}
