package main

import (
	"fmt"
	"math"
)

/*
Let’s model different shapes (Circle and Rectangle) that
implement a Shape interface with methods for calculating the area and perimeter.

Both Circle and Rectangle structs implement the Shape interface
by providing their own Area and Perimeter methods.
The printShapeDetails function can now accept any Shape, making the code flexible and reusable.
*/

// define Shape interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle struct
type Circle struct {
	radius float64
}

// Rectangle struct
type Rectangle struct {
	width, height float64
}

// implement Area & Perimeter methods for Circle
func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// implement Area and Perimeter methods for Rectangle
func (r Rectangle) Area() float64 {
	return r.width * r.height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func printShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

func practicalExample() {
	fmt.Println("\nPractical Example: Using Interfaces to Build Shapes")

	circle := Circle{radius: 5}
	rectangle := Rectangle{width: 4, height: 6}

	// print details of the Circle
	fmt.Println("Circle:")
	printShapeDetails(circle)

	// print details of the Rectangle
	fmt.Println("\nRectangle:")
	printShapeDetails(rectangle)
}
