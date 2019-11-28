package main

import (
	"testing"
	)

func TestSquare(t *testing.T) {
	var s Figure = Square{side:6}

	var expectedArea float64 = 36
	if s.area() != expectedArea {
		t.Errorf("Area was incorrect, got: %v, want: %v.", s.area(), expectedArea)
	}

	var expectedPerimeter float64 = 24
	if s.perimeter() != expectedPerimeter {
		t.Errorf("Perimeter was incorrect, got: %v, want: %v.", s.perimeter(), expectedPerimeter)
	}
}

func TestCircle(t *testing.T) {
	var c Figure = Circle{r:6}

	var expectedArea float64 = 113.09733552923255
	if c.area() != expectedArea {
		t.Errorf("Area was incorrect, got: %v, want: %v.", c.area(), expectedArea)
	}

	var expectedPerimeter float64 = 37.69911184307752
	if c.perimeter() != expectedPerimeter {
		t.Errorf("Perimeter was incorrect, got: %v, want: %v.", c.perimeter(), expectedPerimeter)
	}
}