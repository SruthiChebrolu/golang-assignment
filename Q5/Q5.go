package main

import "fmt"

func worker(jobs <-chan int, results chan<- int) {
	for job := range jobs {
		results <- job * 2
	}
}
func main() {
	jobs := make(chan int, 1)
	results := make(chan int, 1)

	go worker(jobs, results)
	jobs <- 2
	close(jobs)

	result := <-results
	fmt.Println("Result:", result)

}
