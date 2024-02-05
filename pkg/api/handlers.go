package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"practice/pkg/utils"
)

func HelloWorld(c *fiber.Ctx) error {
	var response = map[string]string{"hello": "world"}
	return c.JSON(response)
}

func GetVenueLineup(c *fiber.Ctx) error {
	cursor, err := utils.GetCollection(c.Params("venue")).Find(utils.CTX, bson.M{})
	if err != nil {
		fmt.Println("Error finding documents")
	}

	defer cursor.Close(utils.CTX)

	var response []map[string]interface{}
	cursor.All(utils.CTX, &response)
	return c.JSON(response)
}

func UpdateLineupV1(c *fiber.Ctx) error {
	var payload interface{}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println("error parsing: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	_, err := utils.GetCollection(c.Params("venue")).InsertOne(utils.CTX, payload)

	if err != nil {
		fmt.Println("INSERT ONE ERROR", err)
	}

	return c.JSON(payload)
}
