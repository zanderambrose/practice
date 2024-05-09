package main

import (
	"fmt"
	"os/exec"
)

func main() {
	startAPIService()
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
