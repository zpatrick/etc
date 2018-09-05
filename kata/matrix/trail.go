package main

import "math/rand"
import tm "github.com/buger/goterm"

const ColumnLength = 50

var runes = []rune{
	'a', 'A', 'b', 'B', 'c', 'C', 'd', 'D', 'e', 'E',
	'f', 'F', 'g', 'G', 'h', 'H', 'i', 'I', 'j', 'J',
	'k', 'K', 'l', 'L', 'm', 'M', 'n', 'N', 'o', 'O',
	'p', 'P', 'q', 'Q', 'r', 'R', 's', 'S', 't', 'T',
	'u', 'U', 'v', 'V', 'w', 'W', 'x', 'X', 'y', 'Y', 'z', 'Z',
	'0', '1', '2', '3', '4', '5', '6', '7', '8', '9'}

type Trail struct {
	column  int
	index   int
	content []rune
}

func NewTrail(column int) *Trail {
	content := make([]rune, ColumnLength)
	for i := range content {
		content[i] = runes[rand.Intn(len(runes))]
	}

	return &Trail{
		column:  column,
		content: content,
	}
}

func (t *Trail) Print() {
	for i := 1; i < len(t.content); i++ {
		tm.MoveCursor(t.column, i)
		tm.Printf(string(t.content[i]))
	}
}
