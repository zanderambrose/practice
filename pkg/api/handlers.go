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

func HelloWorld(c *fiber.Ctx) error {
	var response = map[string]string{"hello": "world"}
	return c.JSON(response)
}

func InitClient(client *mongo.Client) *MongoClient {
	return &MongoClient{
		MongoClient: client,
	}
}

func (client *MongoClient) GetVenueLineup(c *fiber.Ctx) error {
	cursor, err := utils.GetCollection(client.MongoClient, c.Params("venue")).Find(utils.CTX, bson.D{})
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

func (client *MongoClient) UpdateLineupV1(c *fiber.Ctx) error {
	var payload map[string]string

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println("error parsing: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	// fmt.Println("payload: ", payload)
	// fmt.Println("params: ", c.Params("venue"))

	id, err := utils.GetCollection(client.MongoClient, c.Params("venue")).InsertOne(utils.CTX, payload)
	fmt.Println("id: ", *id)
	if err != nil {
		fmt.Println("INSERT ONE ERROR", err)
	}

	return c.JSON(payload)
}
