package src

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

// FetchPosts fetches and displays posts from the API
func FetchPosts(baseURL string) {
	resp, err := http.Get(baseURL + "/posts")
	if err != nil {
		fmt.Printf("Error fetching posts: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var posts []Post
	err = json.Unmarshal(body, &posts)
	if err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	for _, post := range posts {
		fmt.Printf("ID: %d, Title: %s, Body: %s\n", post.ID, post.Title, post.Body)
	}
}
