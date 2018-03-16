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
		for i := range tc.sizes {
			if m.Sizes[i] != tc.sizes[i] {
				t.Errorf("expected size %v at position %v, but got %v", tc.sizes[i], i, m.Sizes[i])
			}
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

func TestMLPMats(t *testing.T) {
	tt := []struct {
		name        string
		sizes       []int
		activations []mlp.Activation
		init        mlp.Initializer
	}{
		{"10 hidden", []int{5, 10, 2}, []mlp.Activation{mlp.Identity, mlp.Identity}, mlp.Zeros},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			m, err := mlp.New(
				tc.sizes,
				tc.activations,
				tc.init,
			)
			if err != nil {
				t.Fatalf("unexpected error in New: %v", err)
			}
			d := m.Mats()
			if len(d) != len(tc.sizes)-1 {
				t.Errorf("expected %v weight mats, but got %v", len(tc.sizes)-1, len(d))
			}
			for i := range d {
				r, c := d[i].Dims()
				if r != tc.sizes[i] {
					t.Errorf("expected mat %v to have %v rows, but it had %v", i, tc.sizes[i], r)
				}
				if c != tc.sizes[i+1] {
					t.Errorf("expected mat %v to have %v columns, but it had %v", i, tc.sizes[i+1], c)
				}
			}
		})
	}
}

func TestMLPForward(t *testing.T) {

}

func TestMLPCopy(t *testing.T) {
	m, err := mlp.New(
		[]int{5, 10, 5},
		[]mlp.Activation{mlp.Identity, mlp.Identity},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	c := m.Copy()
	if c == m {
		t.Error("pointer address of copy is the same as source")
	}
	for i := range m.Sizes {
		if c.Sizes[i] != m.Sizes[i] {
			t.Errorf(
				"expected copied size %v at position %v, but got %v",
				m.Sizes[i], i, c.Sizes[i],
			)
		}
	}
	for i := range m.Weights {
		if c.Weights[i] != m.Weights[i] {
			t.Errorf(
				"expected copied weight %v at position %v, but got %v",
				m.Weights[i], i, c.Weights[i],
			)
		}
	}
}
