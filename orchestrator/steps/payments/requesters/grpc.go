package requesters

import (
	"context"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/payments/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type grpcRequester struct {
	port string
}

func NewPaymentsGrpcRequester(port string) Requester {
	return grpcRequester{
		port: ":" + port,
	}
}

func (r grpcRequester) Process(payment *models.Payment) error {
	// Set up a connection to the server.
	conn, err := r.retryDial()
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
	if err != nil {
		return err
	}

	payment.Status = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Revert(payment *models.Payment) error {
	// Set up a connection to the server.
	conn, err := r.retryDial()
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
	if err != nil {
		return err
	}

	payment.Status = convertFromGrpcStatus(res.Status)

	return err
}

func (r grpcRequester) Name() string {
	return "grpc"
}

func convertFromGrpcStatus(grpcStatus proto.PaymentStatus) models.PaymentStatus {
	switch grpcStatus {
	case proto.PaymentStatus_Approved:
		return models.PaymentApproved
	case proto.PaymentStatus_Rejected:
		return models.PaymentRejected
	}
	return models.PaymentNone
}

// grpc settings
var (
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

func (r grpcRequester) retryDial() (*grpc.ClientConn, error) {
	return grpc.Dial(r.port, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithDefaultServiceConfig(retryPolicy))
}
