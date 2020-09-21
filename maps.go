package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["Rob"] = 39
	m["Tania"] = 40
	m["Sylvie"] = 35

	fmt.Println(m)

	sde := m["Rob"]
	fmt.Println(sde)

	delete(m, "Sylvie")
	fmt.Println(m)

	m2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(m2)
}
