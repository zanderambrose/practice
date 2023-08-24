package main

import (
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
		if err := client.Disconnect(utils.CTX); err != nil {
			panic(err)
		}
	}()

	api.ApplyRoutes(app, client)

	app.Listen(":8080")
}
