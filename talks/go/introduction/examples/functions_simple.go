package main

import (
	"fmt"
)

func describeSong(title string, year int) string {
	return fmt.Sprintf("'%s' came out in %d.", title, year)
}

func main() {
	fmt.Println("Welcome to the official Hoobastank fan club!")
	fmt.Println(describeSong("The Reason", 2003))
}
