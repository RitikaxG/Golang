package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(input <-chan int, output chan<- int, wg *sync.WaitGroup) { // Takes 2 channels as input ( loops through value of input channel { receive only channel } ) sends its square to output channel { send only channel }
	defer wg.Done()
	for val := range input { // Loops through values in input channel
		time.Sleep(time.Second)
		output <- val * 2 // Returns square of values to output channel
	}

}

func WorkerPool() {
	// Create 2 channels
	input := make(chan int)
	output := make(chan int)

	var wg sync.WaitGroup

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go worker(input, output, &wg) // Spawns 2 goroutine to get the square of values in input channel concurrently
	}

	go func() {
		defer close(input) // After all values are sent close input channel
		for i := 0; i < 5; i++ {
			input <- i // Sends i to input channel
		}
	}()

	go func() {
		wg.Wait() // As soon as all goroutines are done u should close output channel
		close(output)
	}()

	// for i := 0; i < 5; i++ {
	// 	println(<-output) // Instead of looping through channel u explicitly expected exactly 5 results, so even though u never closed output channel no deadlock occurred
	// }

	for val := range output { // Now that output channel is closed u can range over it, its not waiting for more values..
		fmt.Println(val)
	}

}
