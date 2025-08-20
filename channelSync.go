package main

import "fmt"

// How to ensure that main thread exists only when the goroutine has executed

// 1. Add Unbuffered channel that block the receiver

func ExpensiveSum(ch chan bool) {
	sum := 0
	for i := 1; i < 1000000; i++ {
		sum += i
	}
	ch <- true // Send bool to channel

}

func Synchro() {
	ch := make(chan bool)
	go ExpensiveSum(ch)
	<-ch // Recieve ( This will ensure main fn is executed only when expensiveSum operation is completed )
	fmt.Println("Main fn ended")
}

// Sending sum value
func ExpensiveSum2(ch chan int) {
	sum := 0
	for i := 1; i < 1000000; i++ {
		sum += i
	}
	ch <- sum
} // Here we do not use defer cause the go routine exists naturally after sending sum, no resources are left hanging

func Synchro2() {
	ch := make(chan int)
	go ExpensiveSum2(ch)
	sum := <-ch
	fmt.Println(sum)
}
