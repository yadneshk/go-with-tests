package datastructures

import (
	"testing"
)

func TestArea(t *testing.T) {

	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}

	t.Run("rectangles", func(t *testing.T) {
		rect := Rectangle{Width: 10.0, Height: 10.0}
		checkArea(t, rect, 100.0)

	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{Radius: 10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}

func TestPerimeter(t *testing.T) {
	perimeterTests := []struct {
		name         string
		shape        Shape
		expectedArea float64
	}{
		{name: "Rectange", shape: Rectangle{Width: 10, Height: 10.0}, expectedArea: 40.0},
		{name: "Circle", shape: Circle{Radius: 10.0}, expectedArea: 62.83185307179586},
		{name: "Triange", shape: Triange{A: 10, B: 20, C: 30}, expectedArea: 60.0},
	}

	for _, tt := range perimeterTests {
		got := tt.shape.Perimeter()
		if got != tt.expectedArea {
			t.Errorf("%#v got %g want %g", tt.shape, got, tt.expectedArea)
		}
	}
}
