package main

import "fmt"

// START OMIT

func Speak() {
	fmt.Println("Hello, World!")
}

func IsEven(n int) bool {
	return n%2 == 0
}

// END OMIT

func SpeakName(name string) {
	fmt.Printf("%s says: Hello, World!", name)
}

func SpeakNamePhrase(name, phrase string) {
	fmt.Printf("%s says: %s", name, phrase)
}

func SpeakNamePhraseN(name, phrase string, times int) {
	for i := 0; i < times; i++ {
		SpeakPhrase(name, phrase)
	}
}

func main() {}
