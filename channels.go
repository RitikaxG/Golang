package main

// Safely share data betweeen goroutines and aviod race condition

import (
	"fmt"
	"time"
)

func UnbufferedChannels() {
	// Creating a new channel
	ch := make(chan int)
	ch2 := make(chan int)

	go func() { // Store data in channel from goroutine
		ch <- 4
	}()

	go func() {
		ch2 <- 6
	}()
	value := <-ch // Receive data from channel in main thread
	value2 := <-ch2

	fmt.Println("Value stored in channel", value)
	fmt.Println("Value stored in second channel", value2)
}

// Storing multiple values in a single channel
func MultipleValues() {
	ch := make(chan int)
	go func() {
		for i := 1; i < 4; i++ {
			ch <- i
		}
		close(ch)
	}()

	for val := range ch {
		fmt.Println(val)
	}

}

// Buffered Channel : Has Capacity Specified
func BufferedChannel() {
	ch := make(chan int, 3)
	go func() {
		defer close(ch) // Executed after the function completes
		for i := 0; i < 5; i++ {
			ch <- i
		}
	}()
	time.Sleep(time.Second * 2)
	// for value := range ch {
	// 	fmt.Println(value)
	// }
}

func SingleValue() {
	ch := make(chan int, 1)
	go func() {
		defer close(ch)
		ch <- 4
	}()

	value, ok := <-ch
	fmt.Println(value, ok) // ok : if the value is received
}
