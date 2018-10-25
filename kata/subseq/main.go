package main

import "fmt"

func Subseq(s1, s2 []rune) []rune {
	return recurse(s1, s2, len(s1), len(s2))
}

func recurse(s1, s2 []rune, i, j int) []rune {
	if len(s1) == 0 || len(s2) == 0 {
		return []rune{}
	}

	if s1[i-1] == s2[j-1] {
		return append([]rune{s1[i-1]}, recurse(s1, s2, i-1, j-1)...)
	}

	r1, r2 := recurse(s1, s2, i-1, j), recurse(s1, s2, i, j-1)
	if len(r1) >= len(r2) {
		return r1
	}

	return r2
}

func main() {
	fmt.Println(string(Subseq([]rune("AGGTAB"), []rune("GXTXAYB"))))
}
