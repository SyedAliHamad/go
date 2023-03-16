package main

import (
	"log"
)

// This function recursively sorts the slice using bubble sort
func bubbleSort(arr *[]int, n int) {
	len := len(*arr)
	log.Println(len, " , ", n)
	if (len-n) == 1 || (len-n) == 0 {
		return
	}

	for i := 0; i < len-1-n; i++ {
		if (*arr)[i] > (*arr)[i+1] {
			temp := (*arr)[i]
			(*arr)[i] = (*arr)[i+1]
			(*arr)[i+1] = temp
		}
	}

	bubbleSort(arr, n+1)
	return
}

func main() {
	a := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	bubbleSort(&a, 0)
	log.Println(a)
}
