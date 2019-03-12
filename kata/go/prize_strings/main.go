package main

import (
	"fmt"
	"strings"
)

// https://projecteuler.net/problem=191

// Prize = < 2 Ls, no more than 3 A in a ro

const (
	Absent = 'A'
	Late   = 'L'
	OnTime = 'O'
)

// all Os,
// all permutations of one L
// all permutations of As

func NumPrizes(period int) []string {
	prizes := []string{}
	prizes = append(prizes, strings.Repeat(string(OnTime), period))

	// all permutations of L
	for i := 0; i < period; i++ {
		prize := strings.Repeat(string(OnTime), period)
		prize = replaceAtIndex(prize, i, Late)
		prizes = append(prizes, prize)
	}

	// all permutations of A
	for numA := 0; i < period; i++ {
		i = number of A
                prize := strings.Repeat(string(OnTime), period)
                prize = replaceAtIndex(prize, i, Late)
                prizes = append(prizes, prize)
        }

	// all permutations of L + A

	return prizes
}

func replaceAtIndex(s string, i int, r rune) string {
	res := []rune(s)
	res[i] = r
	return string(res)
}

func main() {
	prizes := NumPrizes(4)
	fmt.Println(prizes)
	fmt.Println("Length:", len(prizes))
}
