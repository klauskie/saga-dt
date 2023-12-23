package grpc

import (
	"context"
	"github.com/klauskie/saga-dt/inventory/proto"
	"google.golang.org/grpc"
)

type InventoryServer struct {
	proto.UnimplementedInventoryServer
}

func NewServer(grpcServer *grpc.Server) {
	inventoryGrpc := &InventoryServer{}
	proto.RegisterInventoryServer(grpcServer, inventoryGrpc)
}

func (s InventoryServer) Deduct(_ context.Context, req *proto.InventoryRequest) (*proto.InventoryResponse, error) {
	return &proto.InventoryResponse{
		UserId:    req.UserId,
		OrderId:   req.OrderId,
		ProductId: req.ProductId,
		Status:    proto.InventoryStatus_Available,
	}, nil
}

func (s InventoryServer) Add(_ context.Context, req *proto.InventoryRequest) (*proto.InventoryResponse, error) {
	return &proto.InventoryResponse{
		UserId:    req.UserId,
		OrderId:   req.OrderId,
		ProductId: req.ProductId,
		Status:    proto.InventoryStatus_Available,
	}, nil
}
