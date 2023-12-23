package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/inventory"
)

type httpRequester struct{}

func NewInventoryHttpRequester() Requester {
	return httpRequester{}
}

func (r httpRequester) Process(inventory *inventory.Inventory) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Revert(inventory *inventory.Inventory) error {
	// TODO implement
	panic("missing implementation")
}

func (r httpRequester) Name() string {
	return "http"
}
