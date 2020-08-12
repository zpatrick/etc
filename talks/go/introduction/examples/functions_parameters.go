package main

import (
	"fmt"
)

func describeAlbum(title string, year, numTracks int) string {
	return fmt.Sprintf("'%s', which came out in %d, featured %d tracks", title, year, numTracks)
}

func main() {
	fmt.Println("Welcome to the official Hoobastank fan club!")
	fmt.Println(describeAlbum("The Reason", 2003, 12))
}
