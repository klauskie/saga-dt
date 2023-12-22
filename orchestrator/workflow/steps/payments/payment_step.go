package payments

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments/requesters"
)

type step struct {
	payment    Payment
	requester  requesters.Requester
	stepStatus steps.Status
}

func NewStep(r models.Order) steps.Step {
	return &step{
		payment: Payment{
			UserId:   r.UserId,
			OrderId:  r.Id,
			Amount:   r.Amount,
			Currency: "USD",
			Status:   None,
		},
		requester: requesters.RetryRequester{
			DefaultRequester: requesters.NewPaymentsGrpcRequester(),
			MaxRetries:       3,
			Delay:            0,
		},
	}
}

func (s *step) Process() bool {
	err := s.requester.Process(&s.payment)
	if err != nil {
		s.stepStatus = steps.Failed
	}

	switch s.payment.Status {
	case Approved:
		s.stepStatus = steps.Complete
	case Rejected:
		s.stepStatus = steps.Failed
		return false
	}

	return true
}

func (s *step) Revert() bool {
	err := s.requester.Revert(&s.payment)
	if err != nil {
		return false
	}
	return true
}

func (s *step) Status() steps.Status {
	return s.stepStatus
}

func (s *step) Name() string {
	return "payments/grpc" // todo combine with fetcher
}
