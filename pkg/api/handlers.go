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

func (client *MongoClient) GetAllUsers(c *fiber.Ctx) error {
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

func (client *MongoClient) CreateUser(c *fiber.Ctx) error {
	var payload models.Person

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println("error parsing: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	payload.IsAdult = utils.IsAdult(payload.Age)

	id, err := utils.GetCollection(client.MongoClient, "people").InsertOne(utils.CTX, payload)
	fmt.Println("id: ", *id)
	if err != nil {
		fmt.Println("INSERT ONE ERROR", err)
	}

	personMap := map[string]interface{}{
		"name":    payload.Name,
		"age":     payload.Age,
		"isAdult": payload.IsAdult,
		"id":      *id,
	}

	return c.JSON(personMap)
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
