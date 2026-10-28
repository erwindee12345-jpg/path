package main

import (
	"fmt"
	"os"
)

var apiKey = os.Getenv("API_KEY")

func isAuthorized(provided string) bool {
	if apiKey == "" {
		return false
	}
	return provided == apiKey
}

func main() {
	
}