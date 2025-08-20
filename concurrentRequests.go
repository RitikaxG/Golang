package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

// Define person type
type person2 struct {
	UserId int    `json:"userId"`
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func sendRequest(url string, ch chan string) {
	res, err := http.Get(url) // Send GET Request
	if err != nil {
		log.Fatal("Unable to send http request", err)
	}

	defer res.Body.Close() // Close Body to free up resources
	body, err := io.ReadAll(res.Body)

	if err != nil {
		log.Fatal("Unable to read data", err)
	}

	var p person
	json.Unmarshal(body, &p) // Take JSON byte and fill them to struct
	ch <- p.Title            // Send title to channel
}

func ConcurrentRequest() {
	ch := make(chan string)
	// Array of type string, size = 3
	urls := [3]string{"https://jsonplaceholder.typicode.com/posts/1", "https://jsonplaceholder.typicode.com/posts/2", "https://jsonplaceholder.typicode.com/posts/3"}

	for _, url := range urls {
		go sendRequest(url, ch) // Start a goroutine that handles request to each url
	}

	for range urls {
		select { // Whichever channels data is received first print them first run loop for no of urls
		case res := <-ch:
			fmt.Println(res)
		}
	}
}
