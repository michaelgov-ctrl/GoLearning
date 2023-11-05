package main

import "fmt"

func sortInts(I []int) (slice []int) {
	sorted := false
	for !sorted {
		sorted = true
		for i := range I[:len(I)-1] {
			if I[i] > I[i+1] {
				sorted = false
				I[i], I[i+1] = I[i+1], I[i]
			}
		}
	}
	return I
}

func main() {
	ints := []int{5, 1, 4, 2, 3}
	fmt.Println(sortInts(ints))
}
