package main

import "fmt"

func main() {
	var a []string = []string{"Foo"}

	// Append
	a = append(a, "Bar")
	fmt.Println(a) // [Foo Baz]

	// Append Vector
	b := []string{"Baz"}
	a = append(a, b...)
	fmt.Println(a) // [Foo Bar Baz]
}
