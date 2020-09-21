package main

import "fmt"

type rectangle struct {
	width, height int
}

func (r *rectangle) area() int {
	return r.width * r.height
}

func main() {

	r1 := rectangle{2, 4}

	fmt.Println(r1.area())
}
