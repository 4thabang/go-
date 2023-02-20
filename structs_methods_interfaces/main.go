package main

import "math"

type Rectangle struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base, Height float64
}

type Shaper interface {
	Area() float64
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Height + r.Width)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	radius := (c.Radius * c.Radius)
	return math.Pi * radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
