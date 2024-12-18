package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

// FetchUsers fetches and displays users from the API
func FetchUsers(baseURL string) {
	resp, err := http.Get(baseURL + "/users")
	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Username: %s, Email: %s\n", user.ID, user.Name, user.Username, user.Email)
	}
}
