package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	for i, v := range a {
		fmt.Printf("[%d]: %d\n", i, v)
	}

	a = append(a, 4, 5, 6)
	fmt.Println(a)
	b := []int{7, 8, 9}
	ab := append(a, b...)
	fmt.Println(ab)

	// initialize slice with 0-100
	// append to slice
	// remove from slice
	// combine two slices

https: //github.com/golang/go/wiki/SliceTricks
}
