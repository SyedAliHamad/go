package main

import "log"

//This function recursively checks if the given string is palindrome
func checkPalindrome(a string) bool {
	len := len(a)

	if len == 0 || len == 1 {
		return true
	}

	if a[0] == a[len-1] {
		return checkPalindrome(a[1 : len-1])
	}

	return false
}

func main() {
	log.Println(checkPalindrome("abbccbba"))
}
