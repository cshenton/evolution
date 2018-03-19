// Package agent defines an abstract interface for an evolution learning agent
// as well as a handful of concrete implementations of some agents, for example
// as multi-layer perceptron policy networks.
package agent

import (
	"github.com/cshenton/evolution/server"
)

// Agent is an evolutionary learning agent.
type Agent interface {
	// Prediction and Evolution
	Forward(in []float64) (out []float64)
	Perturb(s int64)
	Update(s int64, f float64)

	// Copying and serialisation
	Copy() (a Agent)
	ToProto() (pb server.Agent)
	FromProto(server.Agent)
}
