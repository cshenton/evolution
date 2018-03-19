package mlp

// Initializer is a weight initialization function.
type Initializer func(n int) []float64

// Zeros returns a zero vector of weights.
func Zeros(n int) (w []float64) {
	w = make([]float64, n)
	return w
}

// Normal returns a random normal vector of weights.
func Normal(n int) (w []float64) {
	return
}
