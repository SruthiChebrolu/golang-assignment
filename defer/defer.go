package main

import "fmt"

func main() {
	fmt.Println("Opening a file")
	defer fmt.Println("Closing a file")
	fmt.Println("Reading data from a file")
	defer fmt.Println("GoLang")
}
