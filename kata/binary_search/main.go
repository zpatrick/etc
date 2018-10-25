package main

import "fmt"

func IndexOfRecursive(arr []int, v int) int {
	if len(arr) == 0 {
		return -1
	}

	switch mid := len(arr) / 2; {
	case arr[mid] == v:
		return mid
	case mid < v:
		return IndexOf(arr[:mid], v)
	default:
		return IndexOf(arr[mid:], v)
	}
}

func IndexOf(arr []int, v int) int {
	for pivot := len(arr) / 2; pivot >= 0 && pivot < len(arr); {
		switch current := arr[pivot]; {
		case v > current:
			move := (len(arr) - pivot) / 2
			if move == 0 {
				move = 1
			}

			pivot += move
		case v < current:
			move := pivot / 2
			if move == 0 {
				move = 1
			}

			pivot -= move
		default:
			return pivot
		}
	}
	return -1
}

func main() {
	arr := []int{5, 9, 12, 13, 28, 34, 60}
	fmt.Println(IndexOf(arr, 3))
}
