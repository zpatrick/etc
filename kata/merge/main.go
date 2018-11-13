package main

import (
	"fmt"
)

func Merge(lists ...[]int) []int {
	ret := make([]int, 0)
	ptrs := make([]int, len(lists))

	for {
		fmt.Println(ptrs)
		var min, minptr int
		var isSet bool
		for i := 0; i < len(lists); i++ {
			// remove lists that are done
			if len(lists[i]) <= ptrs[i] {
				fmt.Printf("Removing list: %v\n", lists[i])
				lists = append(lists[:i], lists[i+1:]...)
				ptrs = append(ptrs[:i], ptrs[i+1:]...)
				i--
				continue
			}

			// find the next minimum value
			if !isSet || lists[i][ptrs[i]]  < min {
				isSet = true
				min = lists[i][ptrs[i]]
				minptr = i
			}
		}

		if len(lists) == 0 {
			return ret
		}

		ret = append(ret, min)
		ptrs[minptr]++
	}
}

func main() {
	fmt.Println(Merge([]int{1, 9, 32, 158}, []int{66, 89823}, []int{9, 23, 46, 234, 968, 1024}))
}
