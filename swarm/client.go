package swarm

import (
	"context"

	"github.com/cshenton/openai-evolution/mlp"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// Join will call the bootstrap rpc on the specified server, and
// return a reconstructed MLP.
func Join(address string) (m *mlp.MLP, err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	c := NewSwarmClient(conn)
	out := c.Join(context.Background(), &empty.Empty{})
}
