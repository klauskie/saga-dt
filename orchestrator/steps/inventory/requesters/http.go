package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
)

type httpRequester struct{}

func NewInventoryHttpRequester() Requester {
	return httpRequester{}
}

func (r httpRequester) Process(inventory *models.Inventory) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Revert(inventory *models.Inventory) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Name() string {
	return "http"
}
