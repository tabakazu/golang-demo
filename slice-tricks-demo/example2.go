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

	// Get Max Value From Int Slice
	fmt.Println(getMaxValueFromIntSlice(a)) // 10000

	// Get Min Value From Int Slice
	fmt.Println(getMinValueFromIntSlice(a)) // 5

	// Get Sum Value From Int Slice
	fmt.Println(getSumValueFromIntSlice(a)) // 10105

	// Reversing
	reverseIntSlice(a)
	fmt.Println(a) // [10000 5 100]

	// Get 2 Large Values In Int Slice
	a = append(a, []int{400, 9999, 1, 10001, 10}...)
	fmt.Println(a)                               // [10000 5 100 400 9999 1 10001 10]
	fmt.Println(getTwoLargeValuesIntIntSlice(a)) // [100 10000]
}

func getMaxValueFromIntSlice(slice []int) int {
	var max int
	for _, i := range slice {
		if max < i {
			max = i
		}
	}
	return max
}

func getMinValueFromIntSlice(slice []int) int {
	var min int = 9223372036854775807
	for _, i := range slice {
		if min > i {
			min = i
		}
	}
	return min
}

func getSumValueFromIntSlice(slice []int) int {
	var sum int
	for _, i := range slice {
		sum = sum + i
	}
	return sum
}

func reverseIntSlice(slice []int) {
	for i := len(slice)/2 - 1; i >= 0; i-- {
		opp := len(slice) - 1 - i
		slice[i], slice[opp] = slice[opp], slice[i]
	}
}

// Slice の中で大きな 2 つの値を取得する
func getTwoLargeValuesIntIntSlice(slice []int) []int {
	newSlice := make([]int, 2, 2)
	for _, v := range slice {
		if newSlice[0] < newSlice[1] {
			if newSlice[0] < v {
				newSlice[0] = v
			}
		} else {
			if newSlice[1] < v {
				newSlice[1] = v
			}
		}
	}
	return newSlice
}
