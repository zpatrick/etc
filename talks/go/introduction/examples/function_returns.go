package main

import "fmt"

// START OMIT
func Sum(n0, n1 int) int {

}

func SpeakPhrase(name, phrase string) {
	fmt.Printf("%s says: %s", name, phrase)
}

func SpeakPhraseN(name, phrase string, times int) {
	for i := 0; i < times; i++ {
		SpeakPhrase(name, phrase)
	}
}

// END OMIT

func main() {}
