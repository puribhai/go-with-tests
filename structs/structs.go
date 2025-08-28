package structs

import "math"

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (rect Rectangle) Perimeter() float64 {
	return 2 * (rect.Width + rect.Height)
}

func (rect Rectangle) Area() float64 {
	return rect.Width * rect.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}
