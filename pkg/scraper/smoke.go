package scraper

import (
	"fmt"
	// "practice/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Data struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Image       string `json:"image"`
}

type Performance struct {
	Show Data `json:"show"`
}

type Response struct {
	Results []Performance `json:"results"`
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

	for idx, performance := range results.Results {
		if idx > 0 {
			break
		}
		fmt.Println("name: ", performance.Show.Name)
		fmt.Println("description: ", performance.Show.Description)
		fmt.Println("image: ", performance.Show.Image)

	}
}
