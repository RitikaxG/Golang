package main

import "fmt"

type shape interface { // Our interface must have 2 fn area(), perimeter
	area() float32
	perimeter() float32
}

type rectangle struct { // Define rectangle
	width  float32
	height float32
}

type square struct { // Define square
	side float32
}

type circle struct { // Define circle
	radius float32
}

// Make our types (sqaure, rectangle, circle ) obey rules defined in interface (i.e it must have area() perimeter methods )
func (r rectangle) area() float32 { // Fn that returns area of rectangle return type - float32
	return r.width * r.height
}

func (r rectangle) perimeter() float32 { // Fn that returns perimeter of rectangle return type - float32
	return 2 * (r.width * r.height)
}

func (s square) area() float32 {
	return s.side * s.side
}

func (s square) perimeter() float32 {
	return 4 * s.side
}

func (c circle) area() float32 {
	return 3.14 * c.radius * c.radius
}

func (c circle) perimeter() float32 {
	return 2 * 3.14 * c.radius
}

func print(sh shape) { // Define Generic Fn that takes shape as input
	fmt.Println("area", sh.area())
	fmt.Println("perimeter", sh.perimeter())
}

func Interfaces() {
	// Initialise square
	s := square{side: 4}
	r := rectangle{width: 10, height: 12}
	c := circle{radius: 2}

	fmt.Println("Square :")
	print(s)
	fmt.Println("Rectangle :")
	print(r)
	fmt.Println("Circle :")
	print(c)

}
