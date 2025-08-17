package main

import "fmt"

func Variables() {
	var a int = 42
	var b float32 = 3.14
	var c string = "hello"
	var d bool = true

	fmt.Println("Integer", a)
	fmt.Println("Float32", b)
	fmt.Println("String", c)
	fmt.Println("Boolean", d)
}

// Short Hand Syntax
func ShorthandVariables() {
	a := 42
	b := 3.14
	c := "Hello"
	d := true

	fmt.Println("Integer", a)
	fmt.Println("Float", b)
	fmt.Println("String", c)
	fmt.Println("Boolean", d)
}

func Constants() {
	const a int = 42
	const b float32 = 3.14
	const c string = "hello, go"
	const d bool = false

	fmt.Println("Integer", a)
	fmt.Println("Float32", b)
	fmt.Println("String", c)
	fmt.Println("Boolean", d)
}
