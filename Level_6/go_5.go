package main

import (
	"fmt"
	"math"
)

type square struct {
	squareVal float64
}

type circle struct {
	circleVal float64
}

func (s square) area() float64 {
	return s.squareVal * s.squareVal
}

func (c circle) area() float64 {
	return math.Pi * c.circleVal * c.circleVal
}

type shape interface {
	area() float64
}

func info(sh shape) {
	fmt.Println(sh.area())
}

func main() {
	cir := circle{
		circleVal: 12.345,
	}

	sqr := square{
		squareVal: 15,
	}

	info(cir)

	info(sqr)
}
