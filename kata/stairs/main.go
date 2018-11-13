package main

import "fmt"

// You are climbing a stair case.
// It takes n steps to reach to the top.
// Each time you can either climb 1 or 2 steps.
// Create a function that returns each distinct way you climb to the top

func Steps(n int) [][]int {
	return steps([]int{}, n)
}

func steps(current []int, n int) [][]int {
	switch n {
	case 0:
		return [][]int{current}
	case 1:
		return steps(append(current, 1), 0)
	default:
		ret := steps(append(current, 1), n-1)
		return append(ret, steps(append(current, 2), n-2)...)
	}
}

func main() {
	fmt.Println(Steps(4))
}
