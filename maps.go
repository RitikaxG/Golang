package main

import "fmt"

// Key-Value Pair

// It returned false because in Go, A nil map means the map has not been initialized (e.g., var m map[string]int).
// When you use make(map[string]int), Go allocates and initializes the map’s internal data structure — so it’s not nil (even though it’s empty).

// Creates Empty Maps
func CreateMaps() {
	m := make(map[string]int) // Create map with string keys and int values
	fmt.Println(m == nil)
}

// Create and initialise with data
// Output : map[string]int
func CreateInitialiseMaps() {
	m := map[string]int{
		"Aarav": 22,
		"Vansh": 19,
	}
	fmt.Printf("%T\n", m) // Print only the type → you can use fmt.Printf with %T:
	delete(m, "Vansh")    // Delete a key
	fmt.Println(m)

	// Check if a key exists in map
	value, exists := m["Aarav"]
	if exists {
		fmt.Print("Value found", value)
	} else {
		fmt.Print("Value not found")
	}
}
