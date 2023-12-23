package main

import (
	"context"
	"github.com/klauskie/saga-dt/payments/server/grpc"
	"log"
)

func main() {
	if err := grpc.RunServer(context.Background(), "4041"); err != nil {
		log.Fatalf("unable to start inventory server: %v", err)
	}
}
