package main

import "fmt"

func main() {
	x := 100
	y := &x
	fmt.Println(x)
	*y = 12 // changing *y chnages the value stored at that address
	fmt.Println(x)
}
