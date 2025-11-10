package main

import (
	"fmt"
	"sync"
)

func printNumbers(wg *sync.WaitGroup) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
func printLetters(wg *sync.WaitGroup) {
	for ch := 'A'; ch <= 'E'; ch++ {
		fmt.Println(string(ch))
	}
	wg.Done()
}
func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go printNumbers(&wg)
	go printLetters(&wg)
	wg.Wait()
	fmt.Println("Both Goroutines finished")
}
