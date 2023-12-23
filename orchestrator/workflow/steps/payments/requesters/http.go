package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"
)

type httpRequester struct{}

func NewPaymentsHttpRequester() Requester {
	return httpRequester{}
}

func (r httpRequester) Process(payment *payments.Payment) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Revert(payment *payments.Payment) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Name() string {
	return "http"
}
