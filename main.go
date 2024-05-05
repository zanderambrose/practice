package main

import (
	"fmt"
	"os/exec"
)

func main() {
	// Start API service
	fmt.Println("Starting API service...")
	startAPIService()

	// Start scraper service
	fmt.Println("Starting Scraper service...")
	startScraperService()
}

func startAPIService() {
	cmd := exec.Command("go", "run", "./api/main.go")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting API service:", err)
	}
}

func startScraperService() {
	cmd := exec.Command("go", "run", "./scraper/main.go")
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error starting Scraper service:", err)
	}
}
