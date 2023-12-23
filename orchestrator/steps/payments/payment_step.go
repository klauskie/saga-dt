package payments

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/steps"
	"github.com/klauskie/saga-dt/orchestrator/steps/payments/requesters"
)

type step struct {
	payment    models.Payment
	requester  requesters.Requester
	stepStatus steps.Status
}

func NewStep(task models.Task) steps.Step {
	return &step{
		payment: models.Payment{
			UserId:   task.UserId,
			OrderId:  task.OrderId,
			Amount:   task.Amount,
			Currency: "USD",
			Status:   models.PaymentNone,
		},
		stepStatus: steps.Pending,
		requester:  requesters.NewPaymentsGrpcRequester("4041"),
	}
}

func (s *step) Process() error {
	err := s.requester.Process(&s.payment)
	if err != nil {
		s.stepStatus = steps.Failed
		return err
	}

	switch s.payment.Status {
	case models.PaymentApproved:
		s.stepStatus = steps.Complete
	case models.PaymentRejected:
		s.stepStatus = steps.Failed
		return fmt.Errorf("payment step process failed")
	}

	return nil
}

func (s *step) Revert() error {
	return s.requester.Revert(&s.payment)
}

func (s *step) Status() steps.Status {
	return s.stepStatus
}

func (s *step) Name() string {
	return fmt.Sprintf("payments/%s", s.requester.Name())
}
