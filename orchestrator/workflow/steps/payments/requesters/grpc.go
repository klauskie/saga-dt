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
	// Set up a connection to the server.
	conn, err := retryDial()
	if err != nil {
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
	// Set up a connection to the server.
	conn, err := retryDial()
	if err != nil {
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

func (r grpcRequester) Name() string {
	return "grpc"
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

// grpc settings
var (
	addr        = ":8089"
	retryPolicy = `{
		"methodConfig": [{
		  "name": [{"service": "grpc.payments"}],
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
