package main

import "log"

// find the sum of an array recursively
func sum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	return arr[0] + sum(arr[1:])
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}

	log.Println(sum(arr))
}
