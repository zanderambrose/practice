package api

import (
	"github.com/gofiber/fiber/v2"
	"whoshittin/api/handlers"
	"whoshittin/api/middleware"
)

func ApplyRoutes(app *fiber.App) {
	app.Use(middleware.KeyAuthMiddleware)
	app.Get("/api/v1/collections", handlers.ListCollections)
	app.Get("/api/v1/:venue", handlers.GetVenueLineup)
	app.Post("/api/v1/:venue", handlers.UpdateLineupV1)
	app.Delete("/api/v1/:venue", handlers.DeleteCollection)
	app.Post("/api/v1/scrape/venues", handlers.ScrapeVenues)
	app.Post("/api/v1/scrape/:venue", handlers.ScrapeVenue)
}
