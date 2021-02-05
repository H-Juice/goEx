package main

import (
	"fmt"
	"log"
)

type Point struct {
	X int
	Y int
}

type Square struct {
	center Point
	length int
}

func (p *Point) Move(dx int, dy int) {
	p.X += dx
	p.Y += dy
}

func (s *Square) Move(dx int, dy int) {
	s.center.Move(dx, dy)
}

func (s *Square) Area() int {
	return s.length * s.length
}

func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("Invalid length")
	}

	square := &Square{
		center: Point{x, y},
		length: length,
	}
	return square, nil
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("Error")
	}
	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
