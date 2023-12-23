package requesters

import (
	"context"
	proto "github.com/klauskie/saga-dt/inventory/proto"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/inventory"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequester struct{}

func NewInventoryGrpcRequester() Requester {
	return grpcRequester{}
}

func (r grpcRequester) Process(inventory *inventory.Inventory) error {
	// Set up a connection to the server.
	conn, err := retryDial()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewInventoryClient(conn)
	res, err := client.Add(context.Background(), &proto.InventoryRequest{
		UserId:    int64(inventory.UserId),
		OrderId:   inventory.OrderId,
		ProductId: int64(inventory.ProductId),
	})

	inventory.OutOfStock = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Revert(inventory *inventory.Inventory) error {
	// Set up a connection to the server.
	conn, err := retryDial()
	if err != nil {
		return err
	}
	defer conn.Close()

	client := proto.NewInventoryClient(conn)
	res, err := client.Deduct(context.Background(), &proto.InventoryRequest{
		UserId:    int64(inventory.UserId),
		OrderId:   inventory.OrderId,
		ProductId: int64(inventory.ProductId),
	})

	inventory.OutOfStock = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Name() string {
	return "grpc"
}

func convertFromGrpcStatus(grpcStatus proto.InventoryStatus) (outOfStock bool) {
	switch grpcStatus {
	case proto.InventoryStatus_Available:
		outOfStock = false
	case proto.InventoryStatus_Unavailable:
		outOfStock = true
	}
	return outOfStock
}

// grpc settings
var (
	addr        = ":8088"
	retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "grpc.inventory"}],
		  "waitForReady": true,
		  "retryPolicy": {
			  "MaxAttempts": 3,
			  "InitialBackoff": ".01s",
			  "MaxBackoff": ".01s",
			  "BackoffMultiplier": 1.0,
			  "RetryableStatusCodes": [ "UNAVAILABLE" ]
		  }
		}]}`
)

func retryDial() (*grpc.ClientConn, error) {
	return grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
}
