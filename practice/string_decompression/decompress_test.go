package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDecompress(t *testing.T) {
	cases := map[string]string{
		"abc2[x]":     "abcxx",
		"2[ab3[c]d]e": "abcccdabcccde",
		"aaaaa":       "aaaaa",
		"5[a]":        "aaaaa",
		"a2[a]2[a]":   "aaaaa",
		"a2[2[a]]":    "aaaaa",
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			output, err := Decompress(input)
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, output, expected)
		})
	}
}

func TestDecompressErrors(t *testing.T) {
	cases := map[string]string{
		"Non-number before bracket": "a[a]",
		"Extra open bracket":        "2[[a]",
		"Extra closed backet":       "2[a]]",
		"Number without bracket":    "2a",
	}

	for name, input := range cases {
		t.Run(name, func(t *testing.T) {
			if _, err := Decompress(input); err == nil {
				t.Fatalf("Error was nil!")
			}
		})
	}
}
