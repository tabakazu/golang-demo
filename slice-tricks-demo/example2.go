package main

import "fmt"

func main() {
	var a []int = []int{100}

	// Append
	a = append(a, 5)
	fmt.Println(a) // [100 5]

	// Append Vector
	b := []int{10000}
	a = append(a, b...)
	fmt.Println(a) // [100 5 10000]

	// Get Max Int From Slice
	fmt.Println(getMaxIntFromSlice(a)) // 10000
}

func getMaxIntFromSlice(slice []int) int {
	var max int
	for _, i := range slice {
		if max < i {
			max = i
		}
	}
	return max
}
