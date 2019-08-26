package main

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

func invalidCharacterError(position int, c rune) error {
	return fmt.Errorf("Invalid character at position %d %s", position, strconv.QuoteRune(c))
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func Decompress(input string) (string, error) {
	stack := []string{}

	for i := 0; i < len(input); i++ {
		c := rune(input[i])
		switch {
		case unicode.IsLetter(c) || c == '[':
			stack = append(stack, string(c))
		case unicode.IsDigit(c):
			// find the index of the end of the number
			j := strings.IndexFunc(input[i:], func(r rune) bool {
				return r == '['
			})

			if j == -1 {
				return "", fmt.Errorf("No open bracket found after number starting at index %d", i)
			}

			number := input[i : i+j]
			if _, err := strconv.Atoi(number); err != nil {
				return "", fmt.Errorf("Non-number used before open bracket starting at index %d", i)
			}

			stack = append(stack, number)
			i += j - 1
		case c == ']':
			var block string
			var matchFound bool

			for len(stack) > 0 {
				top := stack[len(stack)-1]
				stack = stack[:len(stack)-1]

				if top == "[" {
					matchFound = true
					break
				}

				block += top
			}

			if !matchFound {
				return "", fmt.Errorf("No open bracket found for closed bracket at index %d", i)
			}

			number := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			count, err := strconv.Atoi(number)
			if err != nil {
				return "", fmt.Errorf("Non-number used before open bracket")
			}

			block = reverse(block)
			stack = append(stack, strings.Repeat(block, count))
		default:
			return "", invalidCharacterError(i, c)
		}
	}

	var output string
	for _, s := range stack {
		output += s
	}

	return output, nil
}
