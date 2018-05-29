package main

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

type ReplacerFN func(b []byte) ([]byte, error)

func NewPatternByteReplacer(pattern, replacement []byte) ReplacerFN {
	return func(b []byte) ([]byte, error) {
		return bytes.Replace(b, pattern, replacement, -1), nil
	}
}

func NewStaticByteReplacer(replacement []byte) ReplacerFN {
	return func(b []byte) ([]byte, error) {
		return replacement, nil
	}
}

func NewInteractiveReplacer(r io.Reader, w io.Writer, replacer ReplacerFN) ReplacerFN {
	return func(before []byte) ([]byte, error) {
		after, err := replacer(before)
		if err != nil {
			return nil, err
		}

		if bytes.Equal(before, after) {
			return before, nil
		}

		fmt.Fprintf(w, "- %s", before)
		fmt.Fprintf(w, "+ %s", after)
		fmt.Fprintf(w, "replace? (y/n) ")

		var input string
		if _, err := fmt.Fscanln(r, &input); err != nil {
			return nil, fmt.Errorf("Failed to read input : %s", err)
		}

		input = strings.ToLower(input)
		if input == "y" || input == "yes" {
			return after, nil
		}

		return before, nil
	}
}
