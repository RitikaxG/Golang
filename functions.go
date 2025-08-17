package main

func SumFunction(a int, b int) int {
	return a + b
}

// Multiple Return types
func Calculator(a, b int) (int, int) {
	return a + b, a - b
}

// Calling Calculator Function

// sum, sub := Calculator(2, 3)
// fmt.Println(sum, sub)

// Named Return Types
func NamedCalculator(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

// fmt.Println(NamedCalculator(2, 3)) in main

// func main() {

// 	// Anonymous Function
// 	sum, sub := func(a, b int) (sum, sub int) {
// 		sum = a + b
// 		sub = a - b
// 		return
// 	}(2, 3)

// 	fmt.Println(sum, sub)

// }

// Function can take other function as arguement
func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func calculator(a int, b int, fun func(a, b int) int) int {
	return fun(a, b)
}

// Returning Functions ( Function that returns another function )
func multiplier(factor int) func(int) int {
	return func(a int) int {
		return a * factor
	}
}

/*

	double := multiplier(2) // This returns a function that will multiple whatever passed in double by 2
	fmt.Println(double(3))  //  The returned function "double" multiplies 3 by 2

	tripple := multiplier(3)
	fmt.Println(tripple(5))
*/
