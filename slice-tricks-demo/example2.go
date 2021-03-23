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

	// Push
	a = append(a, 200)
	fmt.Println(a) // [100 5 10000 200]

	// Pop
	x, a := a[len(a)-1], a[:len(a)-1]
	fmt.Println(x) // 200
	fmt.Println(a) // [100 5 10000]

	// Get Max Int From Slice
	fmt.Println(getMaxIntFromSlice(a)) // 10000

	// Reversing
	reverseIntSlice(a)
	fmt.Println(a) // [10000 5 100]
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

func reverseIntSlice(slice []int) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}
