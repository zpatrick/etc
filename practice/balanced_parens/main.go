package main

import "fmt"

// can we use a counter instead of a stack? 
func isBalanced(input string) bool {
	var count int
	for i, c := range input {
		// exit early if we know there isn't
		// enough closing brackets to balance
		if count > len(input)-i {
			return false
		}

		switch c {
		case '(':
			count++
		case ')':
			if count == 0 {
				return false
			}

			count--
		}
	}

	return count == 0
}

func main() {
	inputs := []string{
		")",
		"()",
		"hello (world)",
		"hello ((world))",
		"hell(o) wor(ld)",
		"((())",
		"(()))",
		"((()))",
		"()()(())",
		"()())((",
	}

	for _, input := range inputs {
		fmt.Printf("'%s': %v\n", input, isBalanced(input))
	}
}
