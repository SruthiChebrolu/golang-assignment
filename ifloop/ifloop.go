package main

import "fmt"

func main() {
	age := 19
	if age >= 18 {
		fmt.Println("Eligible to vote")
	} else {
		fmt.Println("Not Eligible to vote")
	}
}
