package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Post struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts"

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	var posts []Post
	if err := json.NewDecoder(resp.Body).Decode(&posts); err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	result, err := json.MarshalIndent(posts, "", "  ")
	if err != nil {
		fmt.Println("Error formatting JSON:", err)
		return
	}
	fmt.Println("soal3/main.go:")
	fmt.Println(string(result))
}
