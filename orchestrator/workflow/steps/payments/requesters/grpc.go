package requesters

import (
	"context"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"
	"github.com/klauskie/saga-dt/payments/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequester struct{}

func NewPaymentsGrpcRequester() Requester {
	return grpcRequester{}
}

func (r grpcRequester) Process(payment *payments.Payment) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8089", opts...)
	if err != nil {
		payment.Status = payments.None
		return err
	}
	defer conn.Close()

	client := proto.NewPaymentsClient(conn)

	res, err := client.Debit(context.Background(), &proto.PaymentRequest{
		UserId:   int64(payment.UserId),
		OrderId:  payment.OrderId,
		Amount:   float32(payment.Amount),
		Currency: payment.Currency,
	})

	payment.Status = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Revert(payment *payments.Payment) error {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8089", opts...)
	if err != nil {
		payment.Status = payments.None
		return err
	}
	defer conn.Close()

	client := proto.NewPaymentsClient(conn)

	res, err := client.Credit(context.Background(), &proto.PaymentRequest{
		UserId:   int64(payment.UserId),
		OrderId:  payment.OrderId,
		Amount:   float32(payment.Amount),
		Currency: payment.Currency,
	})

	payment.Status = convertFromGrpcStatus(res.Status)

	return err
}

func convertFromGrpcStatus(grpcStatus proto.PaymentStatus) payments.PaymentStatus {
	switch grpcStatus {
	case proto.PaymentStatus_Approved:
		return payments.Approved
	case proto.PaymentStatus_Rejected:
		return payments.Rejected
	}
	return payments.None
}
