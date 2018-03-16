package mlp_test

import (
	"testing"

	"github.com/cshenton/openai-evolution/mlp"
)

func TestRand(t *testing.T) {
	s := int64(15382)
	n := 100

	a := mlp.Rand(s, n)
	b := mlp.Rand(s, n)

	if len(a) != n {
		t.Fatalf("expected length %v, but it was %v", n, len(a))
	}

	for i := range a {
		if a[i] != b[i] {
			t.Errorf("mismatch at position %v: %v != %v", i, a[i], b[i])
		}
		if a[i] == 0 {
			t.Errorf("unexpected exact zero at position %v", i)
		}
	}
}
