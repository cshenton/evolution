package server

// import (
// 	"context"

// 	"github.com/cshenton/evolution/agent/mlp"
// 	"github.com/golang/protobuf/ptypes/empty"
// )

// // Server implements the grpc SwarmServer interface.
// type Server struct {
// 	result   chan<- *Result
// 	copyReq  chan<- bool
// 	copyResp <-chan *mlp.MLP
// }

// // NewServer creates a new server with nil channels.
// func NewServer() (s *Server) {
// 	s = &Server{
// 		result:   nil,
// 		copyReq:  nil,
// 		copyResp: nil,
// 	}
// 	return s
// }

// // Join requests a copy of the MLP on copyReq, then gets the copy from copyResp.
// func (s *Server) Join(c context.Context, e *empty.Empty) (m *MLP, err error) {
// 	s.copyReq <- true
// 	mlpCopy := <-s.copyResp

// 	sizes := make([]int32, len(mlpCopy.Sizes))
// 	for i := range sizes {
// 		sizes[i] = int32(mlpCopy.Sizes[i])
// 	}

// 	activations := make([]Activation, len(mlpCopy.Activations))
// 	for i := range activations {
// 		activations[i] = Activation(mlpCopy.Activations[i])
// 	}

// 	m = &MLP{
// 		Sizes:       sizes,
// 		Activations: activations,
// 		Weights:     mlpCopy.Weights,
// 	}
// 	return m, nil
// }

// // Notify pushes the result to the result channel.
// func (s *Server) Notify(c context.Context, r *Result) (e *empty.Empty, err error) {
// 	s.result <- r
// 	return &empty.Empty{}, nil
// }
