package requesters

import (
	"context"
	proto "github.com/klauskie/saga-dt/inventory/proto"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequester struct {
	port string
}

func NewInventoryGrpcRequester(port string) Requester {
	return grpcRequester{
		port: ":" + port,
	}
}

func (r grpcRequester) Process(inventory *models.Inventory) error {
	// Set up a connection to the server.
	conn, err := r.retryDial()
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
	if err != nil {
		return err
	}

	inventory.OutOfStock = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Revert(inventory *models.Inventory) error {
	// Set up a connection to the server.
	conn, err := r.retryDial()
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
	if err != nil {
		return err
	}

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

func (r grpcRequester) retryDial() (*grpc.ClientConn, error) {
	return grpc.Dial(r.port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
}
