package main

import (
	"fmt"
	"math"
)

type square struct {
	edgeLength float64
}
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (s square) area() float64 {
	return s.edgeLength * s.edgeLength
}

type shape interface {
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

func main() {
	sq := square{20}
	info(sq)
	cir := circle{2}
	info(cir)
}
