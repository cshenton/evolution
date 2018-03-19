// Package mlp implements a simple multilayer perceptron in addition to basic
// interfaces for describing one, such as activations and initializers.
package mlp

import (
	"errors"
	"fmt"

	"gonum.org/v1/gonum/mat"
)

// MLP is a multilayer perceptron.
type MLP struct {
	Sizes       []int
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
		Sizes:       sizes,
		Activations: activations,
		Weights:     init(n),
	}

	return m, nil
}

// Mats constructs the weight matrices from the slice of weights.
func (m *MLP) Mats() (d []*mat.Dense) {
	d = make([]*mat.Dense, len(m.Sizes)-1)
	start := 0
	for i := range d {
		end := start + m.Sizes[i]*m.Sizes[i+1]
		d[i] = mat.NewDense(m.Sizes[i], m.Sizes[i+1], m.Weights[start:end])
		start = end
	}
	return d
}

// activationFuncs returns this MLP's activation funcs.
func (m *MLP) activationFuncs() (f []ActivationFunc) {
	f = make([]ActivationFunc, len(m.Activations))
	for i, a := range m.Activations {
		f[i] = a.Func()
	}
	return f
}

// Forward makes a forward pass through the MLP, sequentially multiplying by
// each weight matrix and applying each activation.
func (m *MLP) Forward(in []float64) (out []float64, err error) {
	if len(in) != m.Sizes[0] {
		err := fmt.Errorf("in must have length %v, but was length %v", m.Sizes[0], len(in))
		return nil, err
	}
	afuncs := m.activationFuncs()
	prev := mat.NewDense(1, len(in), in)
	for i, w := range m.Mats() {
		next := mat.NewDense(1, m.Sizes[i+1], nil)
		next.Mul(prev, w)
		next.Apply(func(i, j int, v float64) float64 { return afuncs[i](v) }, next)
		prev = next
	}

	out = make([]float64, m.Sizes[len(m.Sizes)-1])
	fmt.Println(len(out))
	fmt.Println(prev.Dims())
	for i := range out {
		out[i] = prev.At(0, i)
	}
	return out, nil
}

// Copy makes a deep copy of the target MLP.
func (m *MLP) Copy() (c *MLP) {
	s := make([]int, len(m.Sizes))
	copy(s, m.Sizes)

	a := make([]Activation, len(m.Activations))
	copy(a, m.Activations)

	w := make([]float64, len(m.Weights))
	copy(w, m.Weights)

	c = &MLP{
		Sizes:       s,
		Activations: a,
		Weights:     w,
	}
	return c
}
