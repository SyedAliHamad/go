package main

import "log"

//checks if the aray is sorted recursively
func isSorted(arr []int, size int) bool {

	if size == 1 {
		return true
	}
	if arr[0] > arr[1] {
		return false
	}

	return isSorted(arr[1:], size-1)
}

func main() {
	arr := []int{1, 2, 3, 5, 5, 6}
	log.Println(isSorted(arr, 6))
}
