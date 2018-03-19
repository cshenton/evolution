package main

import (
	"github.com/cshenton/openai-evolution/mlp"
	"github.com/cshenton/openai-evolution/swarm"
	"github.com/cshenton/seer/server"
)

const (
	internalBuff = 10
	externalBuff = 10
	copiesBuff   = 10
)

func toAny(pb Message) (a &any.Any, err error) {
	typeUrl, _ := "github.com/cshenton/evolution/" + proto.MessageName(pb),
	value, _ := proto.Marshal(pb)
	a = &any.Any{
		TypeUrl: typeUrl,
		Value:   value,
	}
	return a
}

func fromAny(a &any.Any) (ag Agent, err error) {

}

func main() {
	// members.Join() ...
	members := []string{"127.0.0.1"}

	internal := make(chan swarm.Result, internalBuff)
	external := make(chan swarm.Result, externalBuff)
	copyReqs := make(chan bool, copiesBuff)
	copyResp := make(chan *mlp.MLP, copiesBuff)

	a := agent.NewMLP() // Can be swapped with any Agent interface
	s := server.New()
	r := runner.New()

	go a.Start(internal, copyReqs, copyResp)
	go s.Start(external, copyReqs, copyResp)

	for {
		select {
		case _ = <-copyReqs:
			copyResp <- agent.Copy()
		case result := <-external:
			agent.Update(result)
		case result := <-internal:
			agent.Update(result)
			for _, m := range members {
				swarm.Notify(1, 1.0, m)
			}
		}
	}
}

// We could have an environment interface, will all synchronous methods
// Then a concrete runner that accepts environments and does fancy stuff
// Similarly, an agent interface, that is all synchronous, and the server that does the fancy
// Internal branching happens at the runner (i.e. multiple envs)
// Runner instantiates on [] of initialised envs
// does its own little round robin

// So main of course only needs the core agent interface methods.
// So too with runner
// But server needs the concrete stuff.
// So server gets an Agent, does a type assertion, goes on with its life.
// Every new agent implementation needs its own server :(
// I really hate that, and it's a real restriction
// Agent data can take any form, and I can't even rely on a nice Marshal method returning []byte
// Since that's not a typed struct like gRPC requires.
// Maybe just have an agent-type enum and send []byte in like msgpack.
// More "correct" would be to use protocol buffer's "any.Any", then the server would
// essentially look up the protocol buffer from a directory of definitions.

// That would be like:
// call rpc Join()
// get "any" message protocol buffer
// have function that looks up concrete type
// serialise into correct concrete type.
// return Agent interface type to the caller
//
// Either way, client method Join() returns Agent
// and server only requires an Agent.

// Agent interface with various impl
// Environment interface with various impl
// Concrete server that talks in Agents
// Concrete runner that talks in Agents and Environments
// Concrete learner that dispatches between the server and the runner (probably just in main)

// Server expects self describing Any proto for bootstrapping
// Each concrete agent has an inline protocol buffer defining its serialised format

// Nice.
