package main

import (
	"context"
	"github.com/klauskie/saga-dt/inventory/server/grpc"
	"log"
)

func main() {
	if err := grpc.RunServer(context.Background(), "4042"); err != nil {
		log.Fatalf("unable to start inventory server: %v", err)
	}
}
