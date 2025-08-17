package main

import "fmt"

// How to ensure that main thread exists only when the goroutine has executed

// 1. Useunbuffered channel

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

func ExpensiveSum2(ch chan int) {
	sum := 0
	for i := 1; i < 1000000; i++ {
		sum += i
	}
	ch <- sum
}

func Synchro2() {
	ch := make(chan int)
	go ExpensiveSum2(ch)
	sum := <-ch
	fmt.Println(sum)
}
