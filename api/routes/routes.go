package api

import (
	"github.com/gofiber/fiber/v2"
	"whoshittin/api/handlers"
)

func ApplyRoutes(app *fiber.App) {
	app.Get("/api/v1/collections", handlers.ListCollections)

	app.Get("/api/v1/:venue", handlers.GetVenueLineup)
	// TODO - This needs auth
	app.Post("/api/v1/:venue", handlers.UpdateLineupV1)
	// TODO - This needs auth
	app.Delete("/api/v1/:venue", handlers.DeleteCollection)

	// TODO - This needs auth
	app.Post("/api/v1/scrape/venues", handlers.ScrapeVenues)
	// TODO - This needs auth
	app.Post("/api/v1/scrape/:venue", handlers.ScrapeVenue)
}
