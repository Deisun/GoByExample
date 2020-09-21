package main

import "fmt"

func main() {

	fmt.Println(add(2, 4))
	fmt.Println(sub(2, 4))

	a, b := multi()

	fmt.Println(a, b)

	nums := []int{2, 5, 7}
	variadic(nums...)

	variadic(44, 2)
}

func add(a int, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

func multi() (int, int) {
	return 3, 7
}

func variadic(nums ...int) {
	for _, num := range nums {
		fmt.Println(num)
	}
}
