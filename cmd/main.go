package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"practice/pkg/models"
	"practice/pkg/utils"
)

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

	app.Post("/api/v1/person", func(c *fiber.Ctx) error {
		var payload models.Person

		if err := c.BodyParser(&payload); err != nil {
			fmt.Println("error parsing: ", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Bad request",
			})
		}

		payload.IsAdult = utils.IsAdult(payload.Age)

		id, err := utils.GetCollection(client, "people").InsertOne(utils.CTX, payload)
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
