package main

import (
	"fmt"
	"os"
)

func main() {
	var input string
	if _, err := fmt.Scan(&input); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	output, err := Decompress(input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(output)
}
