package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// A signal that travels with your request, helps u know if client left, then stop the server
func httpRequest(w http.ResponseWriter, req *http.Request) {
	ctx := req.Context() // Get context from request object

	fmt.Println("Server has started")
	defer fmt.Println("Server has ended")

	select { // You are waiting for either 1. Work finishes after 10s 2. Client cancels
	case <-time.After(time.Second * 10):
		fmt.Fprint(w, "hello")
	case <-ctx.Done(): // When u send a request & close the browser tab before 10s ctx.Error() will be triggered
		err := ctx.Err()
		fmt.Println("Server stopped", err)
		http.Error(w, err.Error(), http.StatusInternalServerError) // w: response writer, err.Error(): converts error to readable format, http.StatusInternalServerError : 500 status code
	}

}

func handleRequest() {
	http.HandleFunc("/hello", httpRequest)
	fmt.Println("Listening on port 8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("Error starting server at port 8080", err)
	}
}
