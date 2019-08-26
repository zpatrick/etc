package main

import "fmt"

// https://adventofcode.com/2017/day/1

func Captcha(digits ...int) int {
	var sum int
	for i := 0; i < len(digits); i++ {
		if i == len(digits)-1 {
			if digits[i] == digits[0] {
				sum += digits[i]
			}

			break
		}

		if digits[i] == digits[i+1] {
			sum += digits[i]
		}
	}

	return sum
}

func main() {
	fmt.Println(Captcha(9, 1, 2, 1, 2, 1, 2, 9))
}
