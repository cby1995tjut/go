package main

import (
	"math"
)

type Circle struct {
	radius float64
}
type Rectangle struct {
	width, height float64
}

func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func (r Rectangle) Area() float64 {
	return r.width * r.height
}
