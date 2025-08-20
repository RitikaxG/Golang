package main

import "fmt"

// Automatically increment iota , initail value = 0

// const (
// 	East = iota
// 	West
// 	North
// 	South
// )

// func Enums() {
// 	fmt.Println(East, West, North, South)
// }

type Direction int

const ( // East, West , North , South has a type of Direction
	East Direction = iota
	West
	North
	South
)

// func Enums() {
// 	fmt.Println(East, West, North, South)
// }

// Creating a method on the direction type
func (d Direction) String() string { // Fn that takes d (of type direction as input) and returns string
	switch d {
	case East:
		return "East"
	case West:
		return "West"
	case North:
		return "North"
	case South:
		return "South"
	default:
		return "Unknow"
	}
}

func Enums() {
	fmt.Println(North)
}

/*

Go has an interface called fmt.Stringer built into its standard library

type Stringer interface { // Just as we define our custom interfaces this is the built in interface of Go
	String() string
}

// Any type in go that has String() string method automatically satisfies this interface

fmt.Println(North)  in Enums() checks does this value implements String() interface, if yes it automatically calls our method instead of printing raw number North 2 , it implements method and returns North

*/
