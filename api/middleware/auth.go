package middleware

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
	"go.mongodb.org/mongo-driver/bson"
	"whoshittin/api/utils"
)

type Client struct {
	ClientId    string     `json:"clientId" bson:"clientId"`
	ApiKey      string     `json:"apiKey" bson:"apiKey"`
	Permissions Permission `json:"permissions" bson:"permissions"`
}

type Permission struct {
	IsAdmin bool `json:"isAdmin" bson:"isAdmin"`
}

var KeyAuthMiddleware = keyauth.New(keyauth.Config{
	ErrorHandler: func(c *fiber.Ctx, err error) error {
		if err == keyauth.ErrMissingOrMalformedAPIKey {
			return c.Status(fiber.StatusUnauthorized).SendString(err.Error())
		}
		if err == InvalidClientId {
			return c.Status(fiber.StatusBadRequest).SendString(err.Error())
		}
		if err == ResourceDenied {
			return c.Status(fiber.StatusForbidden).SendString(err.Error())
		}
		return c.Status(fiber.StatusUnauthorized).SendString("Invalid or expired API Key")
	},
	Validator: validateAPIKey,
})

var ResourceDenied error = errors.New("permission denied for this resource")
var InvalidClientId error = errors.New("invalid client_id")

func queryClient(clientId string) (Client, error) {
	var client Client
	err := db.GetCollection("client").FindOne(db.CTX, bson.M{"clientId": clientId}).Decode(&client)
	if err != nil {
		return Client{}, InvalidClientId
	}

	return client, nil
}

func validateAPIKey(c *fiber.Ctx, token string) (bool, error) {
	httpVerb := c.Method()
	clientId := c.Get("X-Client-ID")

	client, err := queryClient(clientId)
	if err != nil {
		return false, err
	}

	keyFromDb := client.ApiKey
	hashedAPIKey := sha256.Sum256([]byte(keyFromDb))
	hashedKey := sha256.Sum256([]byte(token))
	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) != 1 {
		return false, keyauth.ErrMissingOrMalformedAPIKey
	}
	isAdmin := client.Permissions.IsAdmin

	if !isAdmin && httpVerb != "GET" {
		return false, ResourceDenied
	}

	return true, nil
}
