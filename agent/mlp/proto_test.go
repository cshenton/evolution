package mlp_test

import (
	"testing"

	"github.com/cshenton/evolution/agent/mlp"
)

func TestToFromProto(t *testing.T) {
	m, err := mlp.New(
		[]int{5, 10, 5},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("unexpected error in New: %v", err)
	}

	// Could probably turn the above mlp into a test table
	p, err := m.ToProto()
	if err != nil {
		t.Fatalf("unexpected error in ToProto: %v", err)
	}

	nm := &mlp.MLP{}
	err = nm.FromProto(p)
	if err != nil {
		t.Fatalf("unexpected error in FromProto: %v", err)
	}

	for i := range m.Sizes {
		if m.Sizes[i] != nm.Sizes[i] {
			t.Errorf("expected size %v at position %v, but got %v", m.Sizes[i], i, nm.Sizes[i])
		}
	}

	for i := range m.Activations {
		if m.Activations[i] != nm.Activations[i] {
			t.Errorf(
				"expected activation %v at position %v, but got %v",
				m.Activations[i], i, nm.Activations[i],
			)
		}
	}

	for i := range m.Weights {
		if m.Weights[i] != nm.Weights[i] {
			t.Errorf("expected weight %v at position %v, but got %v", m.Weights[i], i, nm.Weights[i])
		}
	}
}
