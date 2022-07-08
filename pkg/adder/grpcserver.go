package adder

import (
	"context"

	"github.com/MlPablo/simple_GRPC/pkg/api"
)

// GRPCServer ...
type GRPCServer struct{}

func (s *GRPCServer) mustEmbedUnimplementedAdderServer() {
	//TODO implement me
	panic("implement me")
}

// Add ...
func (s *GRPCServer) Add(ctx context.Context, req *api.AddRequest) (*api.AddResponse, error) {
	return &api.AddResponse{Result: req.GetX() + req.GetY()}, nil
}
