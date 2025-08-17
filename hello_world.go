package main // Package Declaration

// Using Multiple Packages
import (
	"fmt"
	"math/rand"
)

func Hello() {
	fmt.Println("hello world")
	fmt.Println("Random number", rand.Int31n(10)) // you pass the upper limit 'n' and randomly generated no. is in [0,n) here 0 to 9
}

// 1. Go compiler expects you to compile your code properly, or else it won't compile.
