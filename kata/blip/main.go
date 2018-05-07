package main

import (
	"fmt"
	"math/rand"
	"time"
)

const NumBlips = 100

func main() {
	blips := make([]*Blip, NumBlips)
	for {
		for i := range blips {
			blips[i] = &Blip{on: rand.Int()%2 == 0}
		}

		for _, blip := range blips {
			fmt.Print(blip.String())
		}

		fmt.Println("")
		time.Sleep(time.Millisecond * 100)
	}
}
