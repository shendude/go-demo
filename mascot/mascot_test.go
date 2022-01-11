package mascot_test

import (
	"testing"

	"github.com/shendude/go-demo/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("Wrong mascot")
	}
}
