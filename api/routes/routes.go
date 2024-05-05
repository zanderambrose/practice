package api

import (
	"github.com/gofiber/fiber/v2"
	"whoshittin/api/handlers"
)

func ApplyRoutes(app *fiber.App) {
	app.Get("/api/v1/", handlers.HelloWorld)
	app.Get("/api/v1/:venue", handlers.GetVenueLineup)
	app.Post("/api/v1/:venue", handlers.UpdateLineupV1)
}
