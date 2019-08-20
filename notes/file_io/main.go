package main

import (
	"log"
	"os"
)

func main() {
	f, err := os.OpenFile("example.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(f)
	log.Printf("Alfa")
	log.Printf("Bravo")
	log.Printf("Charlie")
}
