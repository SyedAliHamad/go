package main

import "log"

//This function recursively find nth power
func nthPower(num int, n int) int {
	if n == 1 {
		return num
	}

	return num * nthPower(num, n-1)
}

func main() {
	log.Println(nthPower(3, 11))
}
