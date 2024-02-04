package api

import (
	"github.com/gofiber/fiber/v2"
)

func ApplyRoutes(app *fiber.App) {
	app.Get("/api/v1/", HelloWorld)
	app.Get("/api/v1/:venue", GetVenueLineup)
	app.Post("/api/v1/:venue", UpdateLineupV1)
}
