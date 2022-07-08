package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/MlPablo/simple_GRPC/pkg/adder"
	"github.com/MlPablo/simple_GRPC/pkg/api"
)

func main() {
	s := grpc.NewServer()
	srv := &adder.GRPCServer{}
	api.RegisterAdderServer(s, srv)

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	if err := s.Serve(l); err != nil {
		log.Fatal(err)
	}
}
