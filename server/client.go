package server

import (
	"context"

	"github.com/cshenton/evolution/agent"
	"github.com/cshenton/evolution/server/pb"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
)

// Join will call the bootstrap rpc on the specified server, and
// then apply the current configuration to the supplied agent.
func Join(a agent.Agent, address string) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c := pb.NewEvolutionClient(conn)
	out, err := c.Join(context.Background(), &empty.Empty{})
	if err != nil {
		return err
	}

	err = a.FromProto(out)
	return err
}

// Notify notifies the specified server of the result of a model evaluation.
func Notify(seed int64, fitness float64, address string) (err error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	c := pb.NewEvolutionClient(conn)

	in := &pb.Result{
		Seed:    seed,
		Fitness: fitness,
	}
	_, err = c.Notify(context.Background(), in)
	return err
}
