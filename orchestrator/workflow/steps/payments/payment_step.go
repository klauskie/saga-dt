package payments

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments/requesters"
)

type step struct {
	payment    Payment
	requester  requesters.Requester
	stepStatus steps.Status
}

func NewStep(task models.Task) steps.Step {
	return &step{
		payment: Payment{
			UserId:   task.UserId,
			OrderId:  task.OrderId,
			Amount:   task.Amount,
			Currency: "USD",
			Status:   None,
		},
		stepStatus: steps.Pending,
		requester:  requesters.NewPaymentsGrpcRequester(),
	}
}

func (s *step) Process() error {
	err := s.requester.Process(&s.payment)
	if err != nil {
		s.stepStatus = steps.Failed
		return err
	}

	switch s.payment.Status {
	case Approved:
		s.stepStatus = steps.Complete
	case Rejected:
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
