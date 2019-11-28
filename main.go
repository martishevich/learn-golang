package main

import (
	"fmt"
	"math"
)

type Figure interface {
	area() float64
	perimeter() float64
}

type Square struct {
	side float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

type Circle struct {
	r float64
}

func (c Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c Circle) perimeter() float64 {
	return math.Pi * c.r * 2
}

func main() {
	var s Figure = Square{side:6}
	var c Figure = Circle{r:6}

	fmt.Println(s.area(), s.perimeter())
	fmt.Println(c.area(), c.perimeter())

}