package main

import "fmt"

func main() {
	s := make([]string, 3)
	fmt.Println(s)

	s[0] = "rob"
	s[1] = "james"
	s[2] = "mike"

	s = append(s, "tania", "ocasio")
	fmt.Println(s)

	c := make([]string, len(s))
	copy(c, s)

	fmt.Println(c)
	fmt.Println(c[1:3])
	fmt.Println(c[:3])
	fmt.Println(c[3:])

	b := []string{"a", "b", "c"}
	fmt.Println(b)

}
