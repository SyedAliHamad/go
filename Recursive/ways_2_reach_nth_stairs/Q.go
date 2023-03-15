package main

import "log"

// You have been given number of stairs. Need to reach nth stair each time you can climb 1 or  stairs.
// retun the number of distinct ways you can climb the stairs
func nthStair(n int) int {

	if n == 0 {
		return 1
	} else if n < 0 {
		return 0
	}

	return nthStair(n-1) + nthStair(n-2)
}

func main() {

	log.Println(nthStair(3))
}
