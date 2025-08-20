package main

import (
	"fmt"
	"sync"
	"time"
)

func mergeChannelValues(chs ...<-chan int) <-chan int { // Fn that takes slice []chs <-chan int (multiple channels ) as input and returns a single read only channel
	mergedCh := make(chan int) // Create a channel to store values from multiple channels
	var wg sync.WaitGroup

	// Loop through all channels ( since slice returns 2 values index, actual value but here we dont care about index so we use _ )
	for _, c := range chs {
		wg.Add(1)
		// Launch a go routine anonymous fn for each channel that takes a channel as input
		go func(c <-chan int) {
			defer wg.Done()
			for value := range c { // Loop through all values stored in the channel
				mergedCh <- value // Pass that value to mergedCh
			}
		}(c) // call this fn immediately pass the arguement c
	}

	// In Go, reading from a channel never stops unless channel is closed , so we need to close mergedCh once all values are sent. This is also to let consumer know no more values are coming
	// If we dont close { for value := range mergedCh } at consumer end will be hung forever leading to goroutine leak (stuck forever)

	go func() {
		wg.Wait()       // Wait for all goroutines to finish sending
		close(mergedCh) // Close mergedCh as soon as all goroutines have finished sending, this will let consumer ( value receiver ) to know all values have come know more values are coming, prevents hung state at consumer end
	}()

	return mergedCh // Early return ( Values can start flowing to consumers as soon as they’re available, not after everything is done. )
	// Once all values have come, mergedCh was closed by a seperate goroutine
}

func FaninFanout() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	// Fan-Out : One job split into multiple workers (goroutines), so they can run in parallel and process faster

	// Worker 1
	go func(c chan int) {
		defer close(c) // After receving all values do not forget to close channel
		for i := 0; i < 4; i++ {
			c <- i
			time.Sleep(time.Millisecond * 500) // Sends values to channel1 with interval of 500ms
		}
	}(ch1)

	// Worker 2
	go func(c chan int) {
		defer close(c)
		for i := 4; i < 8; i++ {
			c <- i
			time.Sleep(time.Millisecond * 300) // Sends values to channel2 with interval of 300ms
		}
	}(ch2)

	// FanIn : Combines result from multiple channels into a single channel ( Merging multiple streams into one )
	mergedCh := mergeChannelValues(ch1, ch2)
	for value := range mergedCh {
		fmt.Println(value)
	}
}

/*
You don’t need select because each goroutine is only tied to one channel.

If you had one goroutine listening to multiple input channels itself, then you’d need select.
*/
