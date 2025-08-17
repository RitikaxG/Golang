package main

import (
	"fmt"
)

func ExpensiveOperation(str string) {
	for i := range str {
		fmt.Println(str, `->`, i)
	}
}

/*
With Goroutines

func main() {
	start := time.Now()

	go ExpensiveOperation("today")
	go ExpensiveOperation("tomorrow")

	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)
	time.Sleep(time.Second * 2) // Sleep for 2 seconds

}


Output :

(base) ritikagupta@ritikas-MacBook-Air-2 Go % go run .
today -> 0
today -> 1
today -> 2
today -> 3
today -> 4
Done in 59.792µs
tomorrow -> 0
tomorrow -> 1
tomorrow -> 2
tomorrow -> 3
tomorrow -> 4
tomorrow -> 5
tomorrow -> 6
tomorrow -> 7

*/

/*
Without Goroutine

func main() {
	start := time.Now()

	ExpensiveOperation("today")
	ExpensiveOperation("tomorrow")

	elapsed := time.Since(start)
	fmt.Println("Done in", elapsed)
	time.Sleep(time.Second * 2) // Sleep for 2 seconds

}

Output:

(base) ritikagupta@ritikas-MacBook-Air-2 Go % go run .
today -> 0
today -> 1
today -> 2
today -> 3
today -> 4
tomorrow -> 0
tomorrow -> 1
tomorrow -> 2
tomorrow -> 3
tomorrow -> 4
tomorrow -> 5
tomorrow -> 6
tomorrow -> 7
Done in 86.667µs


*/
