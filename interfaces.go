package main

import (
	"fmt"
	"math"
)

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

type geometry interface {
	area() float64
	perimeter() float64
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2*r.width + 2*r.height
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println("area", g.area())
	fmt.Println("perimeter:", g.perimeter())
}

func main() {
	r1 := rect{3, 4}
	c1 := circle{3}

	measure(r1)
	measure(c1)
}
