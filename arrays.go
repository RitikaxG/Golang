package main

import "fmt"

// Arrays : Size is fixed at the time of declaration and cant be changed

// 1. Initialise an array of size 5
// Output : [0 0 0 0 0]
func Arrays() {
	var a [5]int
	fmt.Println(a)
}

// 2. Iterate over an array to find some of elements inside
func ArraySum(a [5]int) { // You must define the length of array
	var sum = 0
	for i := 0; i < len(a); i++ {
		sum += a[i]
	}
	fmt.Println(sum)
}

// Calling the function inside main

// arr := [5]int{1, 2, 3, 4, 5}
// ArraySum(arr)

// 3. Find sum of all elements in a 2D Array
func Array2D(a [3][3]int) {
	var sum = 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			sum += a[i][j]
		}
	}
	fmt.Println(sum)
}

// Calling the function inside main

// arr := [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// Array2D(arr)
