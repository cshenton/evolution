package mlp_test

import (
	"testing"

	"github.com/cshenton/openai-evolution/mlp"
)

func TestNew(t *testing.T) {
	tt := []struct {
		name        string
		sizes       []int
		activations []mlp.Activation
		init        mlp.Initializer
		len         int
	}{
		{"10 hidden", []int{5, 10, 2}, []mlp.Activation{mlp.Identity, mlp.Identity}, mlp.Zeros, 70},
	}
	for _, tc := range tt {
		m, err := mlp.New(
			tc.sizes,
			tc.activations,
			tc.init,
		)
		if err != nil {
			t.Fatalf("unexpected error in New: %v", err)
		}
		if m.Input != tc.sizes[0] {
			t.Errorf("expected %v input units, but got %v", tc.sizes[0], m.Input)
		}
		if m.Output != tc.sizes[len(tc.sizes)-1] {
			t.Errorf("expected %v output units, but got %v", tc.sizes[len(tc.sizes)-1], m.Output)
		}
		if len(m.Weights) != tc.len {
			t.Errorf("expected %v weights, but got %v", tc.len, len(m.Weights))
		}
	}
}

func TestNewErrs(t *testing.T) {
	tt := []struct {
		name        string
		sizes       []int
		activations []mlp.Activation
		init        mlp.Initializer
	}{
		{"too few layers", []int{1, 1}, []mlp.Activation{mlp.Identity, mlp.Identity}, mlp.Zeros},
		{"layer, activation mismatch", []int{1, 2, 1}, []mlp.Activation{mlp.Identity}, mlp.Zeros},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m, err := mlp.New(
				tc.sizes,
				tc.activations,
				tc.init,
			)
			if err == nil {
				t.Error("expected error but it was nil")
			}
			if m != nil {
				t.Errorf("expected nil mlp but got %v", m)
			}
		})
	}
}
