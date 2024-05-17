package api

import (
	"crypto/sha256"
	"crypto/subtle"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"whoshittin/api/handlers"
)

func validateAPIKey(c *fiber.Ctx, token string) (bool, error) {
	// Mock db
	apiKeys := map[string]string{
		"client_id1": "api_key",
		"client_id2": "api_key2",
	}

	clientId := c.Get("X-Client-ID")

	dbKey := apiKeys[clientId]
	hashedAPIKey := sha256.Sum256([]byte(dbKey))
	hashedKey := sha256.Sum256([]byte(token))
	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func ApplyRoutes(app *fiber.App) {
	app.Get("/api/v1/collections", handlers.ListCollections)

	app.Use(keyauth.New(keyauth.Config{
		Validator: validateAPIKey,
	}))

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
