package main

import "time"

func worker2(input <-chan int, result chan<- int) {
	println("Entered worker")
	for val := range input {
		time.Sleep(time.Second)

		result <- val * 2
	}
	println("Exited worker")
}

func Puzzle() {
	input := make(chan int)
	results := make(chan int)

	for i := 0; i < 2; i++ {
		go worker2(input, results)
	}

	// Sending values to input channel
	for i := 0; i < 5; i++ {
		input <- i
	}
	close(input)

	// Receiving values from  output channel
	for i := 0; i < 5; i++ {
		println(<-results)
	}

}

/*
Both sending & receiving happens on a single main thread
1st :- Main thread sends all values to input channel
2nd :- As input channel gets populated, worker starts to send square of values to output channel

Both occurs concurrently,
main thread is busy with sending or populating input channel. Meanwhile result channel is getting populated with sqaure of values

But main thread is not able to concurrently receive & print values from output channel as its getting populated cause its busy sending to input channel
Also results channels is never closed so the worker goroutine is always waiting for sending more values to results channel

This leads to DEADLOCK!
so the follwoing code does not run
*/
