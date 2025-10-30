package main

import "fmt"

func main() {
	grade := "A"
	switch grade {
	case "A":
		fmt.Println("Excellent")
	case "B":
		fmt.Println("Good Job")
	case "C":
		fmt.Println("Needs Improvement")
	}
}
