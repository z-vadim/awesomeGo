package main

import "fmt"

func main() {
	mWord := "radar"

	if isPalindrome(mWord) != true {
		fmt.Printf("Word '%v' - isn't palindrome", mWord)
	} else {
		fmt.Printf("Word '%v' - is palindrome", mWord)
	}
}

func isPalindrome(str string) bool {
	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-i-1] {
			return false
		}
	}
	return true
}
