package node

import (
	"bufio"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewListString(t *testing.T) {
	expected := NewNodeString("Alfa", NewNodeString("Bravo", NewNodeString("Charlie", nil)))
	assert.Equal(t, expected, NewListString("Alfa", "Bravo", "Charlie"))
}

func TestScanner(t *testing.T) {
	root := NewListString("Alfa", "Bravo", "Charlie")
	scanner := bufio.NewScanner(root)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		t.Fatal(err)
	}
}
