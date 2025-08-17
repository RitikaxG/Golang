package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// Encoding & Decoding JSON Data
// In Go, only exported fields (those starting with an uppercase letter) are encoded/decoded by the encoding/json package.

// That json:"name" is called a struct tag. It’s metadata you attach to struct fields so the encoding/json package (or other libraries) knows how to map struct fields ↔ JSON keys.
type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func JsonData() {
	person := Person{Name: "ritika", Age: 22, Email: ""}
	jsonData, err := json.Marshal(person)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(jsonData))
}
