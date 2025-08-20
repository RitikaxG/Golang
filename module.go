package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin" // Using external package
)

func UsingGin() {
	// Create Gin Router
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) { // pointer to gin context object that gives access to request data and response writer
		c.JSON(200, gin.H{"message": "pong"}) // gin.H is just a shortcut for map[string]any (a dictionary in Go).
	})

	fmt.Println("Listening on port 8080")
	err := r.Run(":8080")
	if err != nil {
		log.Fatal("Failed to start server at 8080")
	}
}
