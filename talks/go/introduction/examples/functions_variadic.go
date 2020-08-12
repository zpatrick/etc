package main

import (
	"fmt"
)

func printAccomplishments(accomplishments ...string) {
	for _, a := range accomplishments {
		fmt.Printf("* %s\n", a)
	}
}

func main() {
	fmt.Println("Welcome to the official Hoobastank fan club!")
	fmt.Println("As we all know, Hoobastank has won countless awards over the years, including:")
	printAccomplishments(
		"Funniest Band Name - Radio Disney Music Awards",
		"Best International New Artist - NRJ Radio Awards",
		"Video of the Year - Mike awards",
	)
}
