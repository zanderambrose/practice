package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"whoshittin/api/routes"
	"whoshittin/api/utils"
)

func main() {
	app := fiber.New()

	client, err := db.InitDB()

	if err != nil {
		fmt.Println("Failed to connect to MongoDB :", err)
	}

	defer func() {
		if err := client.Disconnect(db.CTX); err != nil {
			panic(err)
		}
	}()

	api.ApplyRoutes(app)

	app.Listen(":8080")
}
