package main

import (
	"fmt"
	"math"
)

func isPowerOfFour(n int) bool {
	if n == 0 {
		return false
	}
	power := (math.Log(float64(n)) / math.Log(4))
	if math.Trunc(power) == power {
		return true
	}
	return false
}

func main() {
	n := []int{16, 1, 2, 5, 7, 8}
	for _, v := range n {
		fmt.Println(isPowerOfFour(v))
	}
}
