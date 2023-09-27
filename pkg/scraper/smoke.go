package scraper

import (
	"fmt"
	// "practice/pkg/utils"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type SmokeData struct {
	BandName          string `json:"bandName"`
	DateOfPerformance string `json:"dateOfPerformance"`
	BandInfo          string `json:"bandInfo"`
}

func Smoke() {
	var data SmokeData
	nodejsServerURL := "http://localhost:3000/scrape"

	// Make an HTTP GET request to the server
	resp, err := http.Get(nodejsServerURL)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error parsing response body: ", err)
		return
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error on unmarshal: ", err)
		return
	}

	formattedJSON, err := json.MarshalIndent(data, "", "    ")
	if err != nil {
		fmt.Println("Error on marshal indent:", err)
		return
	}

	fmt.Println(string(formattedJSON))
}
