package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Struct untuk menyimpan response JSON
type ResponseData struct {
	UserID int    `json:"userId"`
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func fetchDataFromAPI(url string) (*ResponseData, error) {
	// Melakukan HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Membaca response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Parsing JSON ke struct
	var data ResponseData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func main() {
	url := "https://jsonplaceholder.typicode.com/posts/1"
	data, err := fetchDataFromAPI(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Data dari API: %+v\n", data)
}
