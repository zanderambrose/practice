package api

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"go.mongodb.org/mongo-driver/bson"
	"whoshittin/api/handlers"
	"whoshittin/api/utils"
)

type Client struct {
	ClientId    string       `json:"clientId" bson:"clientId"`
	ApiKey      string       `json:"apiKey" bson:"apiKey"`
	Permissions []Permission `json:"permissions" bson:"permissions"`
}

type Permission struct {
	Entry   string `json:"entry" bson:"entry"`
	IsAdmin string `json:"isAdmin" bson:"isAdmin"`
}

func queryClient(clientId string) (Client, error) {
	var client Client
	err := db.GetCollection("client").FindOne(db.CTX, bson.M{"clientId": clientId}).Decode(&client)
	if err != nil {
		fmt.Println("Error finding client")
		return Client{}, errors.New("Error finding client")
	}

	return client, nil
}

func validateAPIKey(c *fiber.Ctx, token string) (bool, error) {
	clientId := c.Get("X-Client-ID")

	client, err := queryClient(clientId)
	if err != nil {
		return false, errors.New("Not a valid client id")
	}

	dbKey := client.ApiKey
	hashedAPIKey := sha256.Sum256([]byte(dbKey))
	hashedKey := sha256.Sum256([]byte(token))
	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}

	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func validateRequest(c *fiber.Ctx, token string) (bool, error) {
	return true, nil
}

func ApplyRoutes(app *fiber.App) {
	app.Get("/api/v1/collections", handlers.ListCollections)

	app.Use(keyauth.New(keyauth.Config{
		Validator: validateRequest,
	}))

	app.Get("/api/v1/:venue", handlers.GetVenueLineup)
	app.Post("/api/v1/:venue", handlers.UpdateLineupV1)
	app.Delete("/api/v1/:venue", handlers.DeleteCollection)
	app.Post("/api/v1/scrape/venues", handlers.ScrapeVenues)
	app.Post("/api/v1/scrape/:venue", handlers.ScrapeVenue)
}
