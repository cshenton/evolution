package mlp_test

import (
	"testing"

	"github.com/cshenton/evolution/agent"
	"github.com/cshenton/evolution/agent/mlp"
)

func TestNew(t *testing.T) {
	tt := []struct {
		name        string
		sizes       []int
		activations []mlp.Activation
		init        mlp.Initializer
		len         int
	}{
		{"10 hidden", []int{5, 10, 2}, []mlp.Activation{mlp.IDENTITY, mlp.IDENTITY}, mlp.Zeros, 70},
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
		{"too few layers", []int{1, 1}, []mlp.Activation{mlp.IDENTITY, mlp.IDENTITY}, mlp.Zeros},
		{"layer, activation mismatch", []int{1, 2, 1}, []mlp.Activation{mlp.IDENTITY}, mlp.Zeros},
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
		{"10 hidden", []int{5, 10, 2}, []mlp.Activation{mlp.IDENTITY, mlp.IDENTITY}, mlp.Zeros},
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
	exp := []float64{0, 0, 0}
	m, err := mlp.New(
		[]int{5, 10, 3},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	out, err := m.Forward([]float64{5, 3, 2, 4, 2}, []int{5})
	if err != nil {
		t.Errorf("unexpected error in Forward: %v", err)
	}

	if len(out) != len(exp) {
		t.Fatalf("expected len %v result, but it was len %v", len(exp), len(out))
	}
	for i := range exp {
		if exp[i] != out[i] {
			t.Errorf("expected %v at position %v, but got %v", exp[i], i, out[i])
		}
	}
}

func TestMLPForwardErrs(t *testing.T) {
	m, err := mlp.New(
		[]int{5, 10, 5},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	out, err := m.Forward([]float64{5, 3, 2, 4}, []int{4})
	if err == nil {
		t.Error("expected error in Forward, but it was nil")
	}
	if out != nil {
		t.Errorf("expected nil output, but it was %v", out)
	}
}

func TestMLPCopy(t *testing.T) {
	m, err := mlp.New(
		[]int{5, 10, 5},
		[]mlp.Activation{mlp.IDENTITY, mlp.IDENTITY},
		mlp.Zeros,
	)
	if err != nil {
		t.Fatalf("expected error in New: %v", err)
	}

	c := m.Copy().(*mlp.MLP) // type assert so we can check the weights, etc.
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

func TestMLPAgent(t *testing.T) {
	// This will force a compilation failure if *mlp.MLP doesn't implement agent.Agent
	_ = agent.Agent(&mlp.MLP{})
}
