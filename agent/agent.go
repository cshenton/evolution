// Package agent defines an abstract interface for an evolution learning agent
// as well as a handful of concrete implementations of some agents, for example
// as multi-layer perceptron policy networks.
package agent

// Agent is an evolutionary learning agent.
type Agent interface {
	Forward(in []float64) (out []float64)
	Perturb(s int64)
	Update(s int64, f float64)
	Copy() (a Agent)
}
