package payments

import (
	"context"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/payments/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

type paymentsGrpcStep struct {
	userId  int
	amount  float64
	orderId string
	status  steps.Status
}

func NewGrpcStep(r models.Order) steps.Step {
	return &paymentsGrpcStep{
		userId:  r.UserId,
		amount:  r.Amount,
		orderId: r.Id,
		status:  steps.Pending,
	}
}

func (s *paymentsGrpcStep) Process() bool {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.Dial(":8089", opts...)
	if err != nil {
		//return err
		log.Fatalf("failed to dial: %v", err)
	}
	defer conn.Close()

	client := proto.NewPaymentsClient(conn)

	res, err := client.Debit(context.Background(), &proto.PaymentRequest{
		UserId:   int64(s.userId),
		OrderId:  s.orderId,
		Amount:   float32(s.amount),
		Currency: "USD",
	})
	if err != nil {
		s.status = steps.Failed
		return false
	}

	switch res.Status {
	case proto.PaymentStatus_Approved:
		s.status = steps.Complete
	case proto.PaymentStatus_Rejected:
		s.status = steps.Failed
		return false
	}

	return true
}

func (s *paymentsGrpcStep) Revert() {}

func (s *paymentsGrpcStep) Status() steps.Status {
	return s.status
}

func (s *paymentsGrpcStep) Name() string {
	return "payments/grpc"
}
