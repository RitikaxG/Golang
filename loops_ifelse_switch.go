package main

import "fmt"

// 1. Find sum of 1st n natural numbers
func Sum(n int) {
	var sum = 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	fmt.Println("Sum", sum)
}

// 2. Range Syntax

// When you use range on a string in Go, the loop variable i is the byte index of each rune in the string.
// "abc" has runes 'a', 'b', 'c' at byte indices 0, 1, and 2.

/* So your loop runs 3 times with:

i = 0
i = 1
i = 2
0 + 1 + 2 = 3

*/

func StringSum() {
	var sum = 0
	for i := range "abc" { //In Go, range works on iterable types: arrays, slices, strings, maps, or channels. int (n) is not iterable, so range n is invalid.
		sum += i
	}
	fmt.Printf("Sum %d\n", sum)
}

func IfElse(i int) {
	if i%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

// Conditional Switch
func SwitchCases(n int) {
	switch {
	case n%2 == 0:
		fmt.Println("even")
	case n%2 != 0:
		fmt.Println("odd")
	default:
		fmt.Println("Neither even nor odd")
	}
}

// Expression Switch
func IdentifyGender(gender string) {
	switch gender {
	case "male":
		fmt.Println("You are a male")
	case "female":
		fmt.Println("You are a female")
	default:
		fmt.Println("You are neither male nor female")
	}
}
