package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) Distance(point *Point) float64 {
	xRes := p.x - point.x
	yRes := p.y - point.y
	return math.Sqrt(math.Pow(xRes, 2) + math.Pow(yRes, 2))
}

func main() {
	p1 := NewPoint(11, 5)
	p2 := NewPoint(0, 0)

	distance := p1.Distance(p2)

	fmt.Println(distance)

}
