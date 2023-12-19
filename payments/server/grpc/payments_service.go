package grpc

import (
	"context"
	"github.com/klauskie/saga-dt/payments/proto"
	"google.golang.org/grpc"
)

type PaymentsServer struct {
	proto.UnimplementedPaymentsServer
}

func NewServer(grpcServer *grpc.Server) {
	paymentsGrpc := &PaymentsServer{}
	proto.RegisterPaymentsServer(grpcServer, paymentsGrpc)
}

func (s PaymentsServer) Debit(_ context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	return &proto.PaymentResponse{
		UserId:  req.UserId,
		OrderId: req.OrderId,
		Amount:  req.Amount,
		Status:  proto.PaymentStatus_Approved,
	}, nil
}

func (s PaymentsServer) Credit(_ context.Context, req *proto.PaymentRequest) (*proto.PaymentResponse, error) {
	return &proto.PaymentResponse{
		UserId:  req.UserId,
		OrderId: req.OrderId,
		Amount:  req.Amount,
		Status:  proto.PaymentStatus_Rejected,
	}, nil
}
