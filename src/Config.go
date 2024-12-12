package src

import (
    "os"
)

type Config struct {
    APIBaseURL string
}

func NewConfig() *Config {
    return &Config{
        APIBaseURL: os.Getenv("API_BASE_URL"),
    }
}