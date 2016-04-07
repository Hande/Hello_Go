package Shape

import (
	"math"
)

type Shape interface {
		Area() float32
		Perimeter() float32
}

type Circle struct {
	Radius float32
}

type Square struct {
	Height float32
	Width float32
}

func (c Circle) Area() float32{
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float32 {
	return math.Pi * c.Radius / 2
}

func (s Square) Area() float32 {
	return s.Height * s.Width
}

func (s Square) Perimeter() float32 {
	return 2 * (s.Height + s.Width)
}