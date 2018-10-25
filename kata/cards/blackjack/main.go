package main

import (
	"fmt"
	"strconv"
)

func RuneToASCII(r rune) string {
	if r < 128 {
		return string(r)
	} else {
		return "\\u" + strconv.FormatInt(int64(r), 16)
	}
}

func main() {
	thing := '♠'
	fmt.Printf("%q", thing)
}
