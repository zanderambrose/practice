package scraper

import (
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"reflect"
	"strings"
	"whoshittin/scraper/utils"
)

type Performer struct {
	Instrument string `json:"instrument"`
	Name       string `json:"name"`
}

type SmallsLiveData struct {
	EventTitle  string      `json:"eventTitle"`
	EventImage  string      `json:"eventImage"`
	EventTime   string      `json:"eventTime"`
	EventDate   string      `json:"eventDate"`
	CurrentTime string      `json:"currentTime"`
	Venue       string      `json:"venue"`
	Band        []Performer `json:"band"`
}

func (data *SmallsLiveData) AppendEventTitle(eventTitle string) {
	data.EventTitle = eventTitle
}

func (data *SmallsLiveData) AppendEventTime(setTime string) {
	data.EventTime = setTime
}

func (data *SmallsLiveData) AppendEventDate(eventDate string) {
	data.EventDate = eventDate
}

func (data *SmallsLiveData) AppendVenue(venue string) {
	venueName := strings.ToLower(venue)
	data.Venue = venueName
}

func (data *SmallsLiveData) AddBandMember(performer Performer) {
	data.Band = append(data.Band, performer)
}

func (data *SmallsLiveData) AppendEventImage(imgSrc string) {
	data.EventImage = "https:" + imgSrc
}

func (data *SmallsLiveData) AppendCurrentTime() {
	time := utils.GetCurrentTime()
	data.CurrentTime = time
}

func SmallsLiveScraper(c *colly.Collector) {
	c.OnHTML("article.event-display-today-and-tomorrow", func(e *colly.HTMLElement) {
		var eventData SmallsLiveData

		ariaLabel := e.ChildAttr("a", "aria-label")
		eventDetails := strings.Split(ariaLabel, ", ")

		// Get event img
		e.ForEach("div.event-picture", func(_ int, elem *colly.HTMLElement) {
			imgSrc := elem.ChildAttr("img", "src")
			eventData.AppendEventImage(imgSrc)
		})

		// Get event title
		eventTitle := e.ChildText("p.event-info-title")
		eventData.AppendEventTitle(eventTitle)

		// Get venue info
		venueInfo, err := getEventDetails(eventDetails, -2)
		if err != nil {
			fmt.Println("Error getting venue information")
		}
		venue := strings.Split(venueInfo, "Live at ")
		eventData.AppendVenue(venue[1])

		// Get set time info
		setTimeInfo, err := getEventDetails(eventDetails, -1)
		if err != nil {
			fmt.Println("Error getting set time information")
		}
		eventTime := strings.Split(setTimeInfo, "sets start at ")
		eventData.AppendEventTime(eventTime[1])

		// Get set date info
		e.ForEach("div.sub-info__date-time", func(_ int, elem *colly.HTMLElement) {
			info := elem.ChildText("div.title5:first-child")
			eventData.AppendEventDate(info)
		})

		// Get performers info
		e.ForEach("div.title5", func(_ int, elem *colly.HTMLElement) {
			var performer Performer
			text := elem.Text
			parts := strings.Split(text, " / ")
			if len(parts) == 2 {
				name := strings.TrimSpace(parts[0])
				instrument := strings.TrimSpace(parts[1])
				performer.Name = name
				performer.Instrument = instrument
				eventData.AddBandMember(performer)
			}
		})

		eventData.AppendCurrentTime()

		// POST data to server
		utils.PostVenueData(venue[1], eventData)

	})

	c.Visit("https://www.smallslive.com/")
}

func getEventDetails(details []string, target int) (string, error) {
	if target == -1 {
		return details[len(details)-1], nil
	}

	if target == -2 {
		return details[len(details)-2], nil
	}

	return "", errors.New("Cannot get index")
}

func printStruct(s interface{}) {
	v := reflect.ValueOf(s)
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		fieldName := t.Field(i).Name
		fieldValue := field.Interface()

		fmt.Printf("%s: %v\n", fieldName, fieldValue)
	}
}
