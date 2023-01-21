package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (s square) area() float64 {
	return s.length * s.length
}

type shape interface {
	// area() - 컴파일 에러
	area() float64
}

func info(s shape) {
	fmt.Println(s.area())
}

var g func() = func() {
	fmt.Println("G")
}
var a func()

func main() {
	x := circle{
		radius: 2.1,
	}
	y := square{
		length: 3,
	}

	info(x)
	info(y)

	g()
	a = g
	a()
}
