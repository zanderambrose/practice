package api

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func ApplyRoutes(app *fiber.App, client *mongo.Client) {
	withClient := InitClient(client)
	app.Get("/api/v1", withClient.GetAllUser)
}
