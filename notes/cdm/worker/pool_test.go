package worker

import (
	"os"
	"testing"

	"github.com/getoutreach/logutil"
)

func TestPool(t *testing.T) {
	p := NewPool(logutil.NewJSONLogger(os.Stdout, logutil.Debug, nil))
	submit, close := p.Listen()
	defer close()

	submit("alfa")
}
