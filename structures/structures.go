package main

import "fmt"

type Address struct {
	city  string
	state string
}

type Employee struct {
	name string
	age  int
}

func main() {
	e1 := Employee{
		name: "sruthi",
		age:  20}
	add1 := Address{
		city:  "Richardson",
		state: "Texas",
	}
	fmt.Println(e1.name)
	fmt.Println(e1.age)
	fmt.Println(add1.city)
	fmt.Println(add1.state)

}
