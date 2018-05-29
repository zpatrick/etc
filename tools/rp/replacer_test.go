package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPatternByteReplacer(t *testing.T) {
	replacer := NewPatternByteReplacer([]byte("foo"), []byte("bar"))
	output, err := replacer([]byte("foo bar baz"))
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "bar bar baz", string(output))
}

func TestInteractiveReplacer(t *testing.T) {
	staticReplacer := NewStaticByteReplacer([]byte("bar"))
	cases := map[string]string{
		"y":   "bar",
		"Y":   "bar",
		"yes": "bar",
		"Yes": "bar",
		"YES": "bar",
		"n":   "foo",
		"N":   "foo",
		"no":  "foo",
		"No":  "foo",
		"NO":  "foo",
		"x":   "foo",
		"sdl": "foo",
	}

	for input, expected := range cases {
		t.Run(input, func(t *testing.T) {
			r := bytes.NewBuffer(nil)
			if _, err := r.WriteString(input); err != nil {
				t.Fatal(err)
			}

			replacer := NewInteractiveReplacer(r, ioutil.Discard, staticReplacer)
			output, err := replacer([]byte("foo"))
			if err != nil {
				t.Fatal(err)
			}

			assert.Equal(t, expected, string(output))
		})
	}
}
