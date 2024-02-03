package scraper

import (
	"fmt"
	// "practice/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type ShowData struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Performance struct {
	Metadata ShowData `json:"metadata"`
}

type Results struct {
	Performance []Performance `json:"show"`
}

type Response struct {
	Results []Results `json:"results"`
}

func Smoke() {
	response, err := http.Get("https://tickets.smokejazz.com/api/performance/?booking=true")
	if err != nil {
		fmt.Println("Error from GET request:", err)
		return
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error from RealAll:", err)
		return
	}

	var results Response
	err = json.Unmarshal(body, &results)
	if err != nil {
		fmt.Println("Error from unmarshal:", err)
		return
	}

	for _, item := range results.Results {
		fmt.Printf("Item: %s\n", item.Performance.Metadata.Name)
		fmt.Printf("Item: %s\n", item.Performance.Metadata.Description)
		fmt.Printf("Item: %s\n", item.Performance.Metadata.Image)
	}
}
