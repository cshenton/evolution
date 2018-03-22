// Package atari contains an environment implementation that communicates with
// a cshenton/atari rpc server to evaluate the provided agent.
package atari

import (
	"context"

	"github.com/cshenton/evolution/agent"
	"github.com/cshenton/evolution/environment/atari/proto"
	"google.golang.org/grpc"
)

// Env is the atari environment.
type Env struct {
	Environment *proto.Environment
	Client      proto.AtariClient
}

// New connects to the specified atari rpc server and returns an Env.
func New(address string, env *proto.Environment) (e *Env, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	e = &Env{
		Environment: env,
		Client:      proto.NewAtariClient(conn)}
	return e, nil
}

// Eval evaluates the provided agent against this environment.
func (e *Env) Eval(a agent.Agent) (f float64) {
	reward := 0.0
	done := false
	o, _ := e.Client.Start(context.Background(), e.Environment)
	obs := make([]float64, len(o.Values))
	for i := range obs {
		obs[i] = float64(o.Values[i])
	}

	for !done {
		action, _ := agent.DiscreteAction(a, obs)
		state, _ := e.Client.Step(context.Background(), &proto.Action{Value: uint32(action)})

		for i := range obs {
			obs[i] = float64(state.Observation.Values[i])
		}
		reward += float64(state.Reward)
		done = state.Done
	}
	return reward
}
