package main

import "fmt"

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) *person {
	p := person{name: name, age: age}
	return &p
}
func main() {

	p1 := newPerson("Rob", 39)

	fmt.Println(p1.name)
	fmt.Println(person{"Tania", 40})

}
