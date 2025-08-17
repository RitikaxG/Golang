package main

import (
	"fmt"
	"sync"
)

// 1. No synchronization - race condition without mutex
func SumRace() {
	var wg sync.WaitGroup
	sum := 0
	for i := 0; i < 10000; i++ {
		wg.Add(1) // A go routine is going to start
		go func() {
			defer wg.Done() // Tells that goroutine has completed its execution runs after fn completes
			sum = sum + 1   // All goroutines accessing the shared data at same time , leading to race condition ( simultaneous updation )
		}()
	}

	wg.Wait() // Waits for all goroutine to finish
	fmt.Println(sum)
}

// 2. Synchronization using Mutex
func SumMutex() {
	sum := 0
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer m.Unlock() // Unlocks shared data after fn completes

			m.Lock() // This ensures only one go routine is able to access sum at a time
			sum = sum + 1
		}()
	}
	wg.Wait()
	fmt.Println(sum)
}

// 3. Synchronization using mutex - SafeCounter

// Encapsulation : shared data & mutex bound together
type SafeCounter struct {
	m     sync.Mutex
	value int32
}

func SafeMutex() {
	var wg sync.WaitGroup

	// Initialise
	s := SafeCounter{
		m:     sync.Mutex{},
		value: 0,
	}

	for i := 0; i < 1000000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			defer s.m.Unlock()

			s.m.Lock()
			s.value++
		}()
	}

	wg.Wait()
	fmt.Println(s.value)
}
