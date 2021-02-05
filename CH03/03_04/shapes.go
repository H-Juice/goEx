package main

import (
	"fmt"
	"math"
)

type Point struct {
	X int
	Y int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

type Square struct {
	length float64
}

func (s *Square) Area() float64 {
	return s.length * s.length
}

type Circle struct {
	Radius float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Phi
}

func sumAreas(shapes []Shape) float64 {
	total := 0.0
	for _, s := range shapes {
		total += s.Area()
	}
	return total
}

type Shape interface {
	Area() float64
}

func main() {
	s := &Square{20}
	c := &Circle{10}

	sh := []Shape{s, c}
	fmt.Println(sumAreas(sh))
}
