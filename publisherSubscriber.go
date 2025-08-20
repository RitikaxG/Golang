// Publisher Subscriber in Go ( Braodcast msgs {same data} to all the subscribers )
package main

import (
	"fmt"
	"sync"
)

func publisher(subscribers []chan string, wg *sync.WaitGroup) { // Takes slice of subscribers ( channels ) as input of type string & waitGroup
	defer wg.Done() // Publisher goroutine is done ( finished )
	var wgMsg sync.WaitGroup
	for i := 0; i < 5; i++ { // Broadcast 5 msgs
		msg := fmt.Sprintf("Message %d", i) // instead of printing in console returns formatted msg to console

		for _, sub := range subscribers { // Loop through all the subscribers
			wgMsg.Add(1)
			go func(sub chan string, msg string) { // Send in a separate goroutine to avoid blocking one slow subscriber

				defer wgMsg.Done()
				sub <- msg // Sends msg to subscriber channel
			}(sub, msg)

		}
		wgMsg.Wait() // Wait for all goroutines to finish
	}

	// Close all subscriber channels once msg broadcasting is done
	for _, sub := range subscribers {
		close(sub)
	}
}

func subscriber(name string, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range c {
		fmt.Println(name, "received", val)
	}
}

func PubSub(subscriberNames []string) {

	var wg sync.WaitGroup
	subscribers := []chan string{} // Create an empty slice to store all subscriber channels
	for _, name := range subscriberNames {
		c := make(chan string)

		subscribers = append(subscribers, c) // Append channel to slice
		wg.Add(1)
		go subscriber(name, c, &wg)
	}

	wg.Add(1)
	go publisher(subscribers, &wg)
	wg.Wait()
}
