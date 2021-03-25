// Interfaces are defined as a set of methods.
// Syntax of interface:
// Type <Interface name> interface {
// <Method name> <Return type>
// }
// In Go, to implement an interface an object just need to implement all methods
// of interface. When an object implement a particular interface, its object can be
// assigned to an interface type variable.

// float64 => 8 bytes
// 1 byte = 8 bit
// float32 => 4 bytes

//concept of Polymorphism is applied
// The word polymorphism means having many forms.
// In simple words, we can define polymorphism as the ability of a
// message to be displayed in more than one form. A real-life example of
// polymorphism, a person at the same time can have different characteristics.
// Like a man at the same time is a father, a husband, an employee.
// So the same person posses different behavior in different situations.
// This is called polymorphism.

//It is an ability of object to behave in multiple form.

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rect struct {
	width  float64
	height float64
}

type Circle struct {
	radius float64
}

func (r Rect) Area() float64 {
	return r.width * r.height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func TotalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.Area()
	}
	return area
}

func TotalPerimeter(shapes ...Shape) float64 {
	var peri float64
	for _, s := range shapes {
		peri += s.Perimeter()
	}
	return peri
}

func main() {
	r := Rect{width: 10, height: 10}
	c := Circle{radius: 10}
	fmt.Println("Total Area: ", TotalArea(r, c))
	fmt.Println("Total Perimeter: ", TotalPerimeter(r, c))
}

// Output:
// Total Area: 414.1592653589793
// Total Perimeter: 102.83185307179586
// Analysis:· A Shape interface is created which contain two methods Area() and
// Perimeter().
// · Rect and Circle implements Shape interface as they implement Area() and
// Perimeter() methods.
// · TotalArea() and TotalPerimeter() are two functions which expect list of
// object of type Shape.
