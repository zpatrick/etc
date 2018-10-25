package main

import "fmt"

// How do you find duplicate numbers in an array if it contains multiple duplicates?

func FindDuplicates(arr []int) []int {
	catalog := map[int]int{}
	for _, key := range arr {
		catalog[key]++
	}

	dups := []int{}
	for key, count := range catalog {
		if count >= 2 {
			dups = append(dups, key)
		}
	}

	return dups
}

// find pairs that add up to sum:
// put them in hash table, see if other exiss
func FindPairs(arr []int, sum int) map[int]int {
	catalog := map[int]bool{}
	for _, key := range arr {
		// this could be map[int]int, and internally we go for i := 0; i < map[key];
		// if we needed to do each sum
		catalog[key] = true
	}

	pairs := map[int]int{}
	for lhs := range catalog {
		rhs := sum - lhs
		if catalog[rhs] {
			pairs[lhs] = rhs
		}
	}

	return pairs
}

func main() {

	fmt.Println(FindDuplicates([]int{1, 2, 3, 4, 5, 2, 8, 5}))
	fmt.Println(FindPairs([]int{1, 2, 3, 4, 5}, 6))
}
