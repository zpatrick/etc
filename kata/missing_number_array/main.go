package main

import "fmt"

// How do you find the missing number in a given integer array of 1 to 100

func FindMissing(arr []int) (int, error) {
	for i := 0; i < len(arr); i++ {
		if arr[i] != i+1 {
			return i, nil
		}
	}

	return 0, fmt.Errorf("No missing integers")
}

func FindMissingRecursive(arr []int) (int, bool) {
	return findMissingRecurse(arr, 1)
}

func findMissingRecurse(arr []int, startingValue int) (int, bool) {
	switch {
	case len(arr) == 0:
		return 0, false
	case len(arr) == 1:
		return startingValue - 1, arr[0] != startingValue
	default:
		middleIndex := len(arr) / 2
		if len(arr)%2 != 0 {
			middleIndex++
		}

		missingValue, found := findMissingRecurse(arr[:middleIndex], startingValue)
		if found {
			return missingValue, true
		}

		return findMissingRecurse(arr[middleIndex:], middleIndex+startingValue)
	}
}

func main() {
	arr := make([]int, 10)
	for i := 0; i < len(arr); i++ {
		arr[i] = i + 1
	}

	arr[8] = 11

	fmt.Println(FindMissing(arr))
	fmt.Println(FindMissingRecursive(arr))
}
