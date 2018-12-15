package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x      int
	y      int
	radius int
}

func (c Circle) Area() float64 {
	return math.Pi * float64(c.radius) * float64(c.radius)
}

func (c Circle) Circumference() float64 {
	return 2 * math.Pi * float64(c.radius)
}

func (c *Circle) Expand() {
	c.radius *= 2
}

func main() {
	var c = Circle{
		x:      100,
		y:      100,
		radius: 50,
	}

	var p = &c

	c.radius = 100 // 先变100
	p.Expand()     // 再*2

	fmt.Printf("%+v\n", c)
	fmt.Printf("%+v\n", p)
}
