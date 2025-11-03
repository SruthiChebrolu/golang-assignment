package main

import "fmt"

func main() {
	nums := [5]int{1, 8, 5, 6, 9}
	max := nums[0]
	for i:=1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println("Maximum number is:", max)
}
