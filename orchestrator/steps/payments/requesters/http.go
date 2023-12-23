package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
)

type httpRequester struct{}

func NewPaymentsHttpRequester() Requester {
	return httpRequester{}
}

func (r httpRequester) Process(payment *models.Payment) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Revert(payment *models.Payment) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Name() string {
	return "http"
}
