package mlp_test

import (
	"testing"

	"github.com/cshenton/openai-evolution/mlp"
)

func TestIdentity(t *testing.T) {
	x := float64(3.1415926)
	if x != mlp.Identity(x) {
		t.Errorf("expected identity to equal %v, but it was %v", x, mlp.Identity(x))
	}
}
