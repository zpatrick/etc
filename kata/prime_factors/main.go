package main

import (
	"fmt"
)

func LargestPrimeFactor(n int) int {
	// find the first factor in n/1, n/2, n/3... that is prime
	for d := 1; d <= n/2; d++ {
		// if d is a factor, find the corresponding int
		if n%d == 0 {
			f := n / d
			if isPrime(f) {
				return f
			}
		}
	}

	// return -1 if there are no prime factors
	return -1
}

func isPrime(n int) bool {
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println(LargestPrimeFactor(600851475143))
}
