package main

type rect struct {
	width  int32
	height int32
}

var r = rect{width: 200, height: 200}

// Function Parameter
func area(r rect) int32 {
	return r.width * r.height
}

// Calling function : fmt.Println(area(r))

// 1. Value receiver
type rect2 struct {
	width  int32
	height int32
}

// Method with receiver (Value Receiver - makes a copy of original struct works on it - performs updation on copied struct )
func (r rect2) area() int32 {
	r.width = 25
	return r.width * r.height
}

// Pointer Reciever (Changes the original structs so changes are persistant outside the method )
func (r *rect2) area2() int32 {
	r.width = 25
	return r.width * r.height
}

/*
func main() {

	r := rect2{40, 40}
	fmt.Println("area", r.area())
	fmt.Println("Value Reciever (original r value initialisation does not change)", r)

	fmt.Println("area", r.area2())
	fmt.Println("Pointer Reciever (r value changes)", r)
	fmt.Println(r.area2())
}

*/

/*

Output :
area 1000
Value Reciever (original r value initialisation does not change) {40 40}
area 1000
Pointer Reciever (r value changes) {25 40}
1000

*/
