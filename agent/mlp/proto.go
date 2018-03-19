package mlp

import (
	mlp_pb "github.com/cshenton/evolution/agent/mlp/pb"
	server_pb "github.com/cshenton/evolution/server/pb"
	"github.com/golang/protobuf/ptypes"
)

// ToProto creates an agent proto from the given MLP.
func (m *MLP) ToProto() (p *server_pb.Agent, err error) {
	a := make([]mlp_pb.Activation, len(m.Activations))
	for i := range a {
		a[i] = mlp_pb.Activation(m.Activations[i])
	}
	s := make([]int32, len(m.Sizes))
	for i := range s {
		s[i] = int32(m.Sizes[i])
	}

	mp := &mlp_pb.MLP{
		Sizes:       s,
		Activations: a,
		Weights:     m.Weights,
	}

	any, err := ptypes.MarshalAny(mp)
	if err != nil {
		return nil, err
	}
	p = &server_pb.Agent{
		Message: any,
	}
	return p, nil
}

// FromProto set's the receiver MLPs weights and config from the provided proto.
func (m *MLP) FromProto(p *server_pb.Agent) (err error) {
	mp := &mlp_pb.MLP{}
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
