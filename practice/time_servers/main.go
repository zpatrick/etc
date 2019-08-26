package main

import (
	"fmt"
	"sort"
)

/*
Given a list of numbers, assume each number represents the amount of time it takes to execute a task.
How would you dive the tasks across two different servers to they finished in the same amount of time?
*/

func Split(tasks []int, n int) [][]int {
	servers := make([][]int, n)
	sums := make([]int, n)
	serverIndex := 0
	max := 0

	sort.Sort(sort.Reverse(sort.IntSlice(tasks)))
	for _, task := range tasks {
		servers[serverIndex] = append(servers[serverIndex], task)
		sums[serverIndex] += task
		if sums[serverIndex] > max {
			max = sums[serverIndex]
			serverIndex++
			if serverIndex == n {
				serverIndex = 0
			}
		}
	}

	return servers
}

func main() {
	tasks := []int{10, 10, 10, 20, 10, 30}
	fmt.Println(Split(tasks, 3))
}
