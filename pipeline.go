package main

import (
	"fmt"
	"math/rand"
)

/*
Pipeline in Go :

Order Pipeline : Step 1 : Takes static data ( slice as input ) to every consecutive step channel is passed as input

*/

// Generates orderId of length 'length' passed by user
func generateOrderId(length int) string {
	chars := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	result := make([]byte, length) // Create a slide of bytes of length 'length'

	// rand.Seed(time.Now().UnixNano()) // This will return time in nanosecond which would be different everytime
	rand.Seed(42)
	for i := 0; i < length; i++ {
		result[i] = chars[rand.Intn(len(chars))] // Picks up a random index from 0 to len(chars) and chars[randomIndex] is assigned to result[i]
	}

	return string(result)
}

type Order struct {
	orderId string
	tax     float32
	price   float32
	name    string
}

// Step 1 : Assign OrderId
func getOrderId(orders []Order) <-chan Order { // Fn that takes slice of orders of type Order as input and returns a read only channel of type Order
	out := make(chan Order) // Create a channel of type Order

	go func() {
		defer close(out) // Close after all orders data has been passed to channel
		for _, o := range orders {
			o.orderId = generateOrderId(16)
			out <- o // Send order to the channel
		}
	}()
	return out // Return channel
}

// Step 2 : Calculate Taxes
func CalculateTaxes(ch <-chan Order) <-chan Order { // Takes channel of type Order as input return channel of type Order as output
	out := make(chan Order)

	go func() {
		defer close(out)
		for o := range ch {
			tax := float32(0)
			if o.price < 1000 {
				tax = 0.1 * o.price
			} else {
				tax = 0.2 * o.price
			}

			o.tax = o.price + tax
			out <- o // After calvulating taxes for orders pass Orders along with taxes to channels
		}
	}()
	return out
}

// Step 3 : Process Orders
func ProcessOrders(ch <-chan Order) {
	for o := range ch {
		fmt.Printf("OrderId : %s | Order Name : %s | Order Price : %.2f | Order Tax : %.2f\n", o.orderId, o.name, o.price, o.tax)
	}
}

// Main Function

func GoOrderPipeline() {
	orders := []Order{
		{name: "Shoes", price: 999},
		{name: "T-Shirt", price: 1999},
	}

	stage1 := getOrderId(orders)
	stage2 := CalculateTaxes(stage1)
	ProcessOrders(stage2)
}
