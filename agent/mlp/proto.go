package mlp

import (
	mlp_proto "github.com/cshenton/evolution/agent/mlp/proto"
	server_proto "github.com/cshenton/evolution/server/proto"
	"github.com/golang/protobuf/ptypes"
)

// ToProto creates an agent proto from the given MLP.
func (m *MLP) ToProto() (p *server_proto.Agent, err error) {
	a := make([]mlp_proto.Activation, len(m.Activations))
	for i := range a {
		a[i] = mlp_proto.Activation(m.Activations[i])
	}
	s := make([]int32, len(m.Sizes))
	for i := range s {
		s[i] = int32(m.Sizes[i])
	}

	mp := &mlp_proto.MLP{
		Sizes:       s,
		Activations: a,
		Weights:     m.Weights,
	}

	any, err := ptypes.MarshalAny(mp)
	if err != nil {
		return nil, err
	}
	p = &server_proto.Agent{
		Message: any,
	}
	return p, nil
}

// FromProto set's the receiver MLPs weights and config from the provided proto.
func (m *MLP) FromProto(p *server_proto.Agent) (err error) {
	mp := &mlp_proto.MLP{}
	err = ptypes.UnmarshalAny(p.Message, mp)
	if err != nil {
		return err
	}

	a := make([]Activation, len(mp.Activations))
	for i := range a {
		a[i] = Activation(mp.Activations[i])
	}
	s := make([]int, len(mp.Sizes))
	for i := range s {
		s[i] = int(mp.Sizes[i])
	}
	m.Sizes = s
	m.Activations = a
	m.Weights = mp.Weights
	return nil
}
