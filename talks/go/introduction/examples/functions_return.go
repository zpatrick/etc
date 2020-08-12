package main

import (
	"errors"
	"fmt"
	"log"
)

// START OMIT
func isGreatestBandToEverExist(band string) (bool, error) {
	if band == "" {
		return false, errors.New("Please specify a band name")
	}

	return band == "Hoobastank", nil
}

// END OMIT

func main() {
	fmt.Println("Welcome to the official Hoobastank fan club!")
	isGreatest, err := isGreatestBandToEverExist("Hoobastank")
	if err != nil {
		log.Fatal(err)
	}

	if !isGreatest {
		log.Fatal("Sorry, something has gone horribly, horribly wrong")
	}

	fmt.Printf("Good news! Hoobastank is still the greatest band to ever exist!")
}
