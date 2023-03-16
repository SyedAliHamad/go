package main

import "log"

// Recursive Search for an element in an aray using binary search
func binarySearch(arr []int, i int) bool {
	len := len(arr)

	if arr[len/2] == i {
		return true
	} else if len == 1 {
		return false
	}

	if arr[len/2] > i {
		return binarySearch(arr[:len/2], i)
	}
	return binarySearch(arr[len/2:], i)
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	log.Println(binarySearch(arr, 11))
}
