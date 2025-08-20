package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func HTTPCient() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1") // Sending a GET request that returns err, response (status code, headers..)
	if err != nil {
		log.Fatal("Unable to send get request", err)
	}

	defer resp.Body.Close() // close Body after getting response to free up resources

	body, err := io.ReadAll(resp.Body) // ReadAll() reads the whole body and returns us []bytes
	if err != nil {
		log.Fatal("Error reading data", err)
	}

	// Convert bytes to string
	fmt.Println(string(body))
}

// Instead of printing raw JSON you can unmarshal it to struct

// Define struct
type person struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func HTTPClient2() {
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		log.Fatal("Unable to send get request", err)
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read data", err)
	}

	// Define a variable of type person
	var p person             // Decoding json into post struct
	json.Unmarshal(body, &p) // Unmarshal needs to modify the original variable (take JSON byte and fill them to struct) so we pass pointer to variable so that it can directly write data into struct

	fmt.Println(p.UserId)
	fmt.Println(p.Body)
	fmt.Println(p.Id)
	fmt.Println(p.Title)

}
