package main

import "log"

// linearly checks if element is present in the array
func linearSearch(arr []int, i int) bool {

	if arr[0] == i {
		return true
	} else if len(arr) == 1 {
		return false
	}
	return linearSearch(arr[1:], i)
}
func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9, 99, 10}
	log.Println(linearSearch(arr, 11))
}
