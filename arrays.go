package main

import "fmt"

func main() {
	var a [5]int
	fmt.Println(a)

	fmt.Println(len(a))

	b := [5]int{2, 4, 6, 8, 10}

	fmt.Println(b)
}
