package main

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReplace(t *testing.T) {
	r := strings.NewReader("foo")
	w := bytes.NewBuffer(nil)
	replacer := NewStaticByteReplacer([]byte("bar"))

	if err := Replace(r, w, replacer); err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, "bar", w.String())
}
