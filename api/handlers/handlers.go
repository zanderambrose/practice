package handlers

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"whoshittin/api/utils"
	"whoshittin/scraper/services"
)

func GetVenueLineup(c *fiber.Ctx) error {
	cursor, err := db.GetCollection(c.Params("venue")).Find(db.CTX, bson.M{})
	if err != nil {
		fmt.Println("Error finding documents")
	}

	defer cursor.Close(db.CTX)

	var response []map[string]interface{}
	cursor.All(db.CTX, &response)
	return c.JSON(response)
}

func UpdateLineupV1(c *fiber.Ctx) error {
	var payload interface{}

	if err := c.BodyParser(&payload); err != nil {
		fmt.Println("error parsing: ", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Bad request",
		})
	}

	_, err := db.GetCollection(c.Params("venue")).InsertOne(db.CTX, payload)

	if err != nil {
		fmt.Println("INSERT ONE ERROR", err)
	}

	return c.JSON(payload)
}

func DeleteCollection(c *fiber.Ctx) error {
	venue := c.Params("venue")
	err := db.GetCollection(venue).Drop(db.CTX)

	if err != nil {
		fmt.Println("DELETE COLLECTION ERROR", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func ListCollections(c *fiber.Ctx) error {
	database := db.GetDatabase()

	collections, err := database.ListCollectionNames(db.CTX, bson.D{{}})
	if err != nil {
		fmt.Println("ERROR GETTING ALL COLLECTIONS:", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	filteredCollections := db.FilterCollectionResults(collections)

	return c.JSON(fiber.Map{
		"collections": filteredCollections,
	})
}

func ScrapeVenue(c *fiber.Ctx) error {
	venue := c.Params("venue")
	scraperFunc := scraper.ScraperMap[venue]

	collector := colly.NewCollector()
	scraperFunc(collector)

	return c.SendStatus(fiber.StatusNoContent)
}

func ScrapeVenues(c *fiber.Ctx) error {
	scraper.Scraper()
	return c.SendStatus(fiber.StatusNoContent)
}
