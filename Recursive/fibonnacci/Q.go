package main

import "log"

// 0,1,1,2,3,5,8,13,21
func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

func main() {
	log.Println(fibonacci(25))
}
