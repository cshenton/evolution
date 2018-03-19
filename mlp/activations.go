package mlp

import "math"

// Activation is a reference to an activation function
type Activation int

const (
	// IDENTITY is the identity function
	IDENTITY Activation = 0
	// TANH is Hyperbolic tan
	TANH Activation = 1
	// RELU is the rectified linear unit
	RELU Activation = 2
)

// Func returns the corresponding activation function.
func (a Activation) Func() ActivationFunc {
	switch a {
	case IDENTITY:
		return Identity
	case TANH:
		return math.Tanh
	case RELU:
		return Relu
	default:
		return Identity
	}
}

// ActivationFunc is a valid activation function.
type ActivationFunc func(x float64) float64

// Identity is the identity activation.
func Identity(x float64) float64 { return x }

// Relu is the rectified linear unit activation.
func Relu(x float64) float64 { return math.Max(0, x) }
