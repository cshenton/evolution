// Package mlp implements a simple multilayer perceptron
package mlp

import (
	"errors"
	"fmt"
)

// MLP is a multilayer perceptron.
type MLP struct {
	Input       int
	Output      int
	Hiddens     []int
	Activations []Activation
	Weights     []float64
}

// New creates a new MLP using the provided configuration, and returns an error
// if there are not as many activations as layers.
func New(sizes []int, activations []Activation, init Initializer) (m *MLP, err error) {
	if len(sizes) < 3 {
		err := errors.New("must have at least 3 layers (input, output and 1 hidden)")
		return nil, err
	}
	if len(activations) != len(sizes)-1 {
		err := fmt.Errorf(
			"must have %v activations for %v hidden layers, but got %v",
			len(sizes)-1, len(sizes)-2, len(activations),
		)
		return nil, err
	}

	n := 0
	for i := range sizes {
		if i == len(sizes)-1 {
			break
		}
		n += sizes[i] * sizes[i+1]
	}

	m = &MLP{
		Input:       sizes[0],
		Output:      sizes[len(sizes)-1],
		Hiddens:     sizes[1 : len(sizes)-1],
		Activations: activations,
		Weights:     init(n),
	}

	return m, nil
}

// Copy makes a deep copy of the target MLP.
func (m *MLP) Copy() (c *MLP) {
	h := make([]int, len(m.Hiddens))
	copy(h, m.Hiddens)

	a := make([]Activation, len(m.Activations))
	copy(a, m.Activations)

	w := make([]float64, len(m.Weights))
	copy(w, m.Weights)

	c = &MLP{
		Input:       m.Input,
		Output:      m.Output,
		Hiddens:     h,
		Activations: a,
		Weights:     w,
	}
	return c
}
