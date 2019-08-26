package main

import (
	"flag"
	"fmt"
)

func chain(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		out <- 1 + <-in
	}()

	return out
}

func main() {
	count := flag.Int("count", 5, "The number of goroutines to use")
	flag.Parse()

	in := make(chan int) 
	out := chain(in)
	for i := 1; i < *count; i++ {
		out = chain(out)
	}
	
	in <- 0
	fmt.Println(<-out)
}
