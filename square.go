package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y int
}

type Square struct {
	start Point
	a     uint
}

func (s Square) End() Point {
	side := int(s.a)
	startPoint := s.start
	return Point{startPoint.x + side, startPoint.y - side}
}

func (s Square) Perimeter() int {
	return int(s.a) * 4
}

func (s Square) Area() int {
	return int(math.Pow(float64(s.a), 2))
}

func main() {
	s := Square{Point{1, 1}, 5}
	fmt.Println(s.End())
	fmt.Println(s.Perimeter())
	fmt.Println(s.Area())
}
