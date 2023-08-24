package main

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"practice/pkg/api"
	"practice/pkg/utils"
)

func main() {
	app := fiber.New()

	client, err := utils.InitDB()

	if err != nil {
		fmt.Println("Failed to connect to MongoDB :", err)
	}

	defer func() {
		if err := client.Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	api.ApplyRoutes(app, client)

	app.Listen(":8080")
}
