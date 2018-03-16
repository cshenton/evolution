package mlp

// Activation is an activation function like tanh or relu.
type Activation func(float64) float64

// Identity is the identity activation.
func Identity(x float64) float64 { return x }
