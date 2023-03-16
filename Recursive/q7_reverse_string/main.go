package main

import (
	"log"
)

// This function recursively reverses the elements in a string
func reverseString(a string) string {
	len := len(a)
	if len == 0 {
		return ""
	} else if len == 1 {
		return a
	}

	first := a[0]
	last := a[len-1]
	a = string(last) + reverseString(a[1:len-1]) + string(first)
	return a
}

func main() {
	str := "lovebabbar"
	log.Println(reverseString(str))
}
