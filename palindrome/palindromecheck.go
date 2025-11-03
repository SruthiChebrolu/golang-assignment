package main

import "fmt"

func main() {
	word := "madam"
	isPalindrome := true

	for i := 0; i < len(word)/2; i++ {
		if word[i] != word[len(word)-1-i] {
			isPalindrome = false
			break
		}
	}
	if isPalindrome {
		fmt.Println(word, "is a palindrome")
	} else {
		fmt.Println(word, "is not a palindrome")
	}
}
