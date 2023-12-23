package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

func RunServer(_ context.Context, port string) error {
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()
	NewServer(grpcServer)

	log.Println("starting grpc payments server...")
	return grpcServer.Serve(lis)
}
