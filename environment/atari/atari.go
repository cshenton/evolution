// Package atari contains an environment implementation that communicates with
// a cshenton/atari rpc server to evaluate the provided agent.
package atari

import "github.com/cshenton/evolution/environment/atari/proto"

// Env is the atari environment.
type Env struct {
	Environment *proto.Environment
	Client   proto.AtariClient
}

// New connects to the specified atari rpc server and returns an Env.
func New(address string, env *proto.Environment) (e *Env, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	e = &Env{
		Environment: env,
		Client: proto.NewAtariClient(conn)}
	return e, nil
}


// Eval evaluates the provided agent against this environment.
func (e *Env) Eval(a agent.Agent) (f float64) {
	reward := 0.0
	done := false
	obs := e.Client.Start(e.Environment)

	for !done {
		action := agent.SampleDiscrete(a.Forward(...obs))
		state := e.client.Step(action)

		obs = state.Observation
		reward += state.Reward
		done = state.Done
	}
	return reward
}