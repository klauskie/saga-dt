package requesters

import "github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"

type Requester interface {
	Process(payment *payments.Payment) error
	Revert(payment *payments.Payment) error
	Name() string
}
