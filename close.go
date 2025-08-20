package main

import "fmt"

// Always close the channel after all the values have been send , or else consumer end will be forever waiting

func CloseChannel() {
	ch := make(chan int)

	go func() {
		defer close(ch) // After fn is finished close channel
		for i := 0; i < 5; i++ {
			ch <- i // Send i to channel
		}
	}()

	for val := range ch { // Loop through channel and print all values in it
		fmt.Println(val)
	}
}
