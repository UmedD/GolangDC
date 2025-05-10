package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
}

type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func main() {
	circle := Circle{radius: 5}
	rectange := Rectangle{Width: 3, Height: 4}

	fmt.Println(circle.Area())
	fmt.Println(rectange.Area())

}
