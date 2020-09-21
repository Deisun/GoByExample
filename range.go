package main

import "fmt"

func main() {

	nums := []int{2, 4, 6}
	sum := 0

	for _, num := range nums {
		sum += num
	}

	fmt.Println(sum)

	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
