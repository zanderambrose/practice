package scraper

import (
	"fmt"
	// "practice/pkg/utils"
	"net/http"
)

func Smoke() {
	data := make(map[string]string, 0)
	nodejsServerURL := "http://localhost:3000/scrape" // Adjust the URL

	// Make an HTTP GET request to the server
	resp, err := http.Get(nodejsServerURL)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	fmt.Println("response: ", resp)

	fmt.Println("data: ", data)
}
