package server

import (
	"context"

	"github.com/cshenton/evolution/agent"
	"github.com/cshenton/evolution/environment"
	"github.com/cshenton/evolution/server/proto"
	"github.com/golang/protobuf/ptypes/empty"
)

// Server implements the grpc SwarmServer interface.
type Server struct {
	result   chan<- environment.Result
	copyReq  chan<- bool
	copyResp <-chan agent.Agent
}

// NewServer creates a new server with nil channels.
func NewServer() (s *Server) {
	s = &Server{
		result:   nil,
		copyReq:  nil,
		copyResp: nil,
	}
	return s
}

// Join requests a copy of the current Agent, then returns a protocol buffer
// representing it.
func (s *Server) Join(c context.Context, e *empty.Empty) (a *proto.Agent, err error) {
	s.copyReq <- true
	ag := <-s.copyResp
	a, err = ag.ToProto()
	return a, err
}

// Notify pushes the result to the result channel.
func (s *Server) Notify(c context.Context, r *proto.Result) (e *empty.Empty, err error) {
	s.result <- environment.Result{
		Seed:    r.Seed,
		Fitness: r.Fitness,
	}
	return &empty.Empty{}, nil
}
