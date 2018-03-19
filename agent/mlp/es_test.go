package mlp_test

import (
	"math"
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

func TestMLPPerturb(t *testing.T) {
	seed := int64(3141)

	m, err := mlp.New(
		[]int{5, 10, 3},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	m.Perturb(seed)
	for i, w := range m.Weights {
		if w == 0 {
			t.Errorf("weight at position %v was not changed", i)
		}
	}
}

func TestMLPUpdate(t *testing.T) {
	seed := int64(3141)
	fitness := 25.0
	r := mlp.Rand(seed, 80)

	m, err := mlp.New(
		[]int{5, 10, 3},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	m.Update(seed, fitness)

	for i, w := range m.Weights {
		if math.Abs(w) > math.Abs(r[i]) {
			t.Errorf("update magnitude %v larger than rand %v at position %v", w, r[i], i)
		}
	}
}
