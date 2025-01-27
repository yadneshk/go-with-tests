package datastructures

import "math"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triange struct {
	A, B, C float64
}

func (rect Rectangle) Area() float64 {
	return rect.Height * rect.Width
}

func (circle Circle) Area() float64 {
	return math.Pi * circle.Radius * circle.Radius
}

func (circle Circle) Perimeter() float64 {
	return math.Pi * 2 * circle.Radius
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Height + rect.Width)
}

func (tri Triange) Perimeter() float64 {
	return tri.A + tri.B + tri.C
}

func (tri Triange) Area() float64 {
	return tri.A + tri.B + tri.C
}
