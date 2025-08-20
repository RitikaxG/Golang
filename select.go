package main

import (
	"fmt"
	"time"
)

func Select() {
	// Create 2 channels
	ch1 := make(chan int)
	ch2 := make(chan int)

	// 1st Goroutine sends data to channel1 after 2s
	go func() {
		time.Sleep(time.Second * 2)
		ch1 <- 2
	}()

	// 2nd goroutine sends data to channel2 after 1s
	go func() {
		time.Sleep(time.Second * 1)
		ch2 <- 6
	}()

	// In the main thread
	for i := 0; i < 2; i++ { // We are running for loop twice because we expect to get 2 values from 2 diff goroutines

		select { // Select monitors which channels data is recieved first, you do not need to block value1:= <- ch1 then proceed to value2 := <-ch2
		// whichever channels data is recieved first that gets printed then for loop runs again for 2nd data and prints it
		case value1 := <-ch1:
			fmt.Println("Value recieved from channel1", value1)
		case value2 := <-ch2:
			fmt.Println("Value received from channel2", value2)
		}

	}
}
