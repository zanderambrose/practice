package main

import (
	"context"
	"fmt"
	// "github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"practice/pkg/utils"
)

type Person struct {
	Name    string `json:"name" bson:"name"`
	Age     int    `json:"age" bson:"age"`
	IsAdult bool   `json:"isAdult" bson:"isAdult"`
}

func isAdult(age int) bool {
	if age < 18 {
		return false
	}
	return true
}

func main() {
	app := fiber.New()

	client, err := utils.InitDB()

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()
	if err != nil {
		fmt.Println("Failed to connect to MongoDB :", err)
	}

	app.Get("/api/v1", func(c *fiber.Ctx) error {
		cursor, err := utils.GetCollection(client, "people").Find(ctx, bson.D{})
		if err != nil {
			fmt.Println("Error finding documents")
		}

		defer cursor.Close(ctx)

		var results []Person
		if err := cursor.All(ctx, &results); err != nil {
			fmt.Println("Error ", err)
		}

		if err := cursor.Err(); err != nil {
			fmt.Println("Error ", err)
		}

		response := map[string]interface{}{
			"people": results,
		}
		return c.JSON(response)
	})

	app.Post("/api/v1/person", func(c *fiber.Ctx) error {
		var payload Person

		if err := c.BodyParser(&payload); err != nil {
			fmt.Println("error parsing: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad request",
			})
		}

		payload.IsAdult = isAdult(payload.Age)

		id, err := getCollection(client, "people").InsertOne(ctx, payload)
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
	})

	app.Delete("/api/v1/:id", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{
			"people": "They are coming",
		})
	})

	app.Listen(":8080")
}
