package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNodeContains(t *testing.T) {
	assert.True(t, NewNodes("alpha", "beta", "charlie").Contains("be"))
}
