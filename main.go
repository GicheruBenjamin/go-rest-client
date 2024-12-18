package main

import (
	"fmt"
	"go-client/src"
)

func main() {
	// Welcome message
	fmt.Println("Welcome to the Go Rest Client!")
	fmt.Println("Choose the entity you wanna see:")
	fmt.Println("----------------------------------")

	// Take user input
	choice := src.Input("Choose the entity 1. Users 2. Posts 3. Comments: ")

	fmt.Println("----------------------------------")

	// Load configuration from .env
	config := src.Config()

	// Handle choices
	switch choice {
	case "1":
		fmt.Println("Fetching users data...")
		src.FetchUsers(config.APIBaseURL)
	case "2":
		fmt.Println("Fetching posts data...")
		src.FetchPosts(config.APIBaseURL)
	case "3":
		fmt.Println("Fetching comments data...")
		src.FetchComments(config.APIBaseURL)
	default:
		fmt.Println("Invalid choice")
	}
	fmt.Println("Thank you for using the Go Rest Client!")
}
