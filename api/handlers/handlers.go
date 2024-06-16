package handlers

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"sync"
	"time"
	"whoshittin/api/utils"
	"whoshittin/scraper/services"
	dateUtils "whoshittin/scraper/utils"
)

var MissingDateParam error = errors.New("missing date param")
var InvalideDateFormat error = errors.New("malformed date param")

func parseDateQueryParam(date string) (string, error) {
	parsedDate, err := time.Parse(dateUtils.STANDARD_DATE_LAYOUT, date)
	if err != nil {
		return "", errors.New("unable to parse date param")
	}
	formattedDate := parsedDate.Format(dateUtils.STANDARD_DATE_REPRESENTATION_LAYOUT)

	return formattedDate, nil
}

func validateDateQueryParam(c *fiber.Ctx) (string, error) {
	dateParam := c.Query("date")

	if len(dateParam) <= 0 {
		return "", MissingDateParam
	}

	date, err := parseDateQueryParam(dateParam)
	if err != nil {
		return "", err
	}

	return date, nil
}

func GetVenueLineup(c *fiber.Ctx) error {
	venue := c.Params("venue")
	date, err := validateDateQueryParam(c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString(err.Error())
	}
	filter := bson.M{"eventDate.formattedDate": date}

	cursor, err := db.GetCollection(venue).Find(db.CTX, filter)
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
	var w sync.WaitGroup
	venue := c.Params("venue")
	scraperFunc := scraper.ScraperMap[venue]

	collector := colly.NewCollector()
	w.Add(1)
	scraperFunc(collector, &w)
	w.Wait()
	return c.SendStatus(fiber.StatusNoContent)
}

func ScrapeVenues(c *fiber.Ctx) error {
	scraper.Scraper()
	return c.SendStatus(fiber.StatusNoContent)
}
