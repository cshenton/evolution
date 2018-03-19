package server

// import (
// 	"context"

// 	"github.com/cshenton/evolution/agent/mlp"
// 	"github.com/golang/protobuf/ptypes/empty"
// 	"google.golang.org/grpc"
// )

// // Join will call the bootstrap rpc on the specified server, and
// // return a reconstructed MLP.
// func Join(address string) (m *mlp.MLP, err error) {
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		return nil, err
// 	}
// 	c := NewSwarmClient(conn)
// 	out, err := c.Join(context.Background(), &empty.Empty{})
// 	if err != nil {
// 		return nil, err
// 	}

// 	s := make([]int, len(out.Sizes))
// 	for i := range s {
// 		s[i] = int(out.Sizes[i])
// 	}
// 	a := make([]mlp.Activation, len(out.Activations))
// 	for i := range a {
// 		a[i] = mlp.Activation(out.Activations[i])
// 	}

// 	m = &mlp.MLP{
// 		Sizes:       s,
// 		Activations: a,
// 		Weights:     out.Weights,
// 	}
// 	return m, nil
// }

// // Notify notifies the specified server of the result of a model evaluation.
// func Notify(seed int64, fitness float64, address string) (err error) {
// 	conn, err := grpc.Dial(address, grpc.WithInsecure())
// 	if err != nil {
// 		return err
// 	}
// 	c := NewSwarmClient(conn)

// 	in := &Result{
// 		Seed:    seed,
// 		Fitness: fitness,
// 	}
// 	_, err = c.Notify(context.Background(), in)
// 	return err
// }
