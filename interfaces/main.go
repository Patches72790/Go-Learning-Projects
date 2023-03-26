package main

import (
	"fmt"
	"math"
)

type Thing interface {
	Do()
}

type Point struct {
	x, y float64
}

func (p *Point) Distance() float64 {
	return math.Abs(p.x - p.y)
}

type Shape struct {
	Point
}

func (s *Shape) Do() {
	fmt.Println("Shape is doing its thing")
}

func DoThings(things []Thing) {
	for _, t := range things {
		t.Do()
	}
}

func main() {
	var things = []Thing{&Shape{Point{x: 1, y: 2}}, &Shape{Point{x: 2, y: 3}}}
	DoThings(things)
}
