package main

import (
	"math/rand"
	"time"

	tm "github.com/buger/goterm"
)

func main() {
	rand.Seed(time.Now().Unix())
	trails := []*Trail{
		NewTrail(0),
	}

	for ; ; time.Sleep(time.Millisecond * 250) {
		tm.Clear()
		for _, trail := range trails {
			trail.Print()
		}

		tm.Flush()
	}
}
