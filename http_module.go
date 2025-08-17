package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.Request contains a lot of values headers, url..passing it by pointer avoid unnecessary copies
func hello(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "hello") // The F stands for "format to". fmt.Fprint(w, "hello") writes "hello" into the writer w, instead of printing to the console
}

func Http() {
	http.HandleFunc("/hello", hello)
	fmt.Println("Listening on port 3000")
	err := http.ListenAndServe(":3001", nil) // nil doesn’t mean "nothing", it means “use the default multiplexer (router)
	if err != nil {
		log.Fatal("Failed to start server")
	}
}
