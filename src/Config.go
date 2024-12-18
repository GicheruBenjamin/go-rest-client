package src

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type config struct {
	APIBaseURL string
}

// Load and validate .env
func Config() config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	apiBaseURL := os.Getenv("API_BASE_URL")
	if apiBaseURL == "" {
		log.Fatal("API_BASE_URL is not set in .env")
	}

	return config{
		APIBaseURL: apiBaseURL,
	}
}
