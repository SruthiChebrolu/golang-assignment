package main

import "fmt"

func main() {
	n := 8 //how many nos to print
	a, b := 0, 1
	//fmt.Println("Fibonacci sequence:")

	for i := 0; i < n; i++ {
		fmt.Println(a, "")
		next_num := a + b
		a = b
		b = next_num
	}
}
