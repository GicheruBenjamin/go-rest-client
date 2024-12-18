package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Comment struct {
	ID     int    `json:"id"`
	Body   string `json:"body"`
	Email  string `json:"email"`
	PostID int    `json:"postId"`
}

// FetchComments fetches and displays comments from the API
func FetchComments(baseURL string) {
	resp, err := http.Get(baseURL + "/comments")
	if err != nil {
		fmt.Printf("Error fetching comments: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var comments []Comment
	err = json.Unmarshal(body, &comments)
	if err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	for _, comment := range comments {
		fmt.Printf("ID: %d, PostID: %d, Email: %s, Body: %s\n", comment.ID, comment.PostID, comment.Email, comment.Body)
	}
}
