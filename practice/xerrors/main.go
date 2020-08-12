package main

import (
	"fmt"
	"log"
	"os"
	"golang.org/x/xerrors"
)


func ReadFile(name string) (string, error) {
	file, err := os.Open(name)
	if err != nil {
		return "", xerrors.Errorf("Failed to read file: %w", err)
	}

	file.Close()
	return "", nil
}

func main() {
	content, err := ReadFile("fake_file")
	if err != nil {
		if pe := new(os.PathError); xerrors.As(err, &pe) {
			log.Fatalf("Could not run operation on %s: %s", pe.Path, err.Error())
		}

		log.Fatalf(err.Error())
	}

	fmt.Println(content)
}
