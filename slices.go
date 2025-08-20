package main

import (
	"fmt"
)

// 1. Create an empty string slice
func CreateStringSlice() {
	var users []string
	fmt.Println(users == nil)
}

// 2. Create and initialise a string slice
func InitialiseStringSlice() {
	var users []string = []string{"asmi", "avni", "ami"}
	fmt.Println(users)
}

// 3. Initialise without default value
func StringSliceInitialise(n int) {
	var users []string = make([]string, n)
	users[0] = "avi"
	users[1] = "ansh"
	users[2] = "abhi"
	// users[3] = "ayaan" this will throw an error if n = 3
	users = append(users, "asmi", "adi", "avni", "ash", "arav", "ami", "ani")
	fmt.Println(users)
}

// Output : [vansh ashmit arav] [vansh ashmit arav]
func CopyByRef() {
	var users []string = []string{"avni", "ashmit", "arav"}
	var users2 = users
	users2[0] = "vansh"
	fmt.Println(users, users2)

}

// Output : [deep raj mrinal] [aarav raj mrinal]
func CopyByValue() {
	var users []string = []string{"deep", "raj", "mrinal"}
	var users2 []string = make([]string, len(users))
	copy(users2, users)
	users2[0] = "aarav"
	fmt.Println(users, users2)
}
