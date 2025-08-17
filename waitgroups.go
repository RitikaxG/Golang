package main

import (
	"fmt"
	"sync"
)

// Allows you to wait for a collection of goroutine to finish execution

// 1. Write a program that runs 5 different go routines and waits for all of them to exit before printing “done”

// a. Without WaitGroups

func ExpensiveOp() {
	str := "ritika"
	for i := range str {
		fmt.Println(i, "->", str)
	}
}

func WaitForChannels(ch chan bool) {
	defer close(ch)
	for i := 0; i < 5; i++ {
		go ExpensiveOp()
		ch <- true
	}
}

func Wait() {
	ch := make(chan bool)
	go WaitForChannels(ch)
	for value := range ch {
		fmt.Println(value) // Sends true into channel right after launching goroutines not waiting for ExpensiveOp to finish
	}
	fmt.Println("Main fn ended")
}

func ExpensiveOp2(ch chan<- bool) { // <- Can only send into channel not receive from it
	defer close(ch)
	str := "ritika"
	for i := range str {
		fmt.Println(i, "->", str)
	}
	ch <- true
}
func SyncedWaits() {
	ch1 := make(chan bool)
	go ExpensiveOp2(ch1)

	ch2 := make(chan bool)
	go ExpensiveOp2(ch2)

	ch3 := make(chan bool)
	go ExpensiveOp2(ch3)

	ch4 := make(chan bool)
	go ExpensiveOp2(ch4)

	ch5 := make(chan bool)
	go ExpensiveOp2(ch5)

	value1 := <-ch1
	fmt.Println(value1)
	value2 := <-ch2
	fmt.Println(value2)
	value3 := <-ch3
	fmt.Println(value3)
	value4 := <-ch4
	fmt.Println(value4)
	value5 := <-ch5
	fmt.Println(value5)

	fmt.Println("Main fn ended")

}

func SyncedWaits2() {
	channels := make([]chan bool, 5)

	for i := 0; i < 5; i++ {
		channels[i] = make(chan bool)
		go ExpensiveOp2(channels[i])
	}

	for i, ch := range channels {
		fmt.Println("Goroutine", i+1, "done", ch)
	}
	fmt.Println("Main fn ended")
}

// b. Waitgroups
func Waitgroups() {
	var wg sync.WaitGroup // Create wait group
	for i := 0; i < 5; i++ {
		wg.Add(1) // Add 1 as loop starts
		go func() {
			defer wg.Done() // Subtract 1 when expensiveOp is executed
			ExpensiveOp()
		}()
	}

	wg.Wait()
	fmt.Println("Main fn ended")
}

func Waitgroups2() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			ExpensiveOp()
		}()
	}
	wg.Wait()
	fmt.Println("DONE")
}
