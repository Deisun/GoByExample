package main

import "fmt"

func main() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	anotherCounter := intSeq()

	fmt.Println(anotherCounter())

}

func intSeq() func() int {
	i := 0

	return func() int {
		i++
		return i
	}
}
