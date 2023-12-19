package main

import (
	"fmt"
	pmtHandler "github.com/klauskie/saga-dt/payments/server/grpc"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	fmt.Println("Payments Service is up...")

	lis, err := net.Listen("tcp", ":8089")
	if err != nil {
		log.Fatalf("cannot initialize server: %v", err)
	}

	grpcServer := grpc.NewServer()
	pmtHandler.NewServer(grpcServer)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalf("impossible to serve: %v", err)
	}
}
