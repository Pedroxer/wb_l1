package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func MakePoint(inx, iny float64) *Point {
	return &Point{
		x: inx,
		y: iny,
	}
}

func (p *Point) X() float64 {
	return p.x
}

func (p *Point) Y() float64 {
	return p.y
}

func main() {
	p1 := MakePoint(1.0, 1.0)
	p2 := MakePoint(5.0, 5.0)
	res := math.Sqrt(math.Pow(p2.X()-p1.X(), 2) + math.Pow(p2.Y()-p1.Y(), 2))
	fmt.Println("res =", res)
}
