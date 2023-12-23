package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/inventory"
)

type Requester interface {
	Process(inventory *inventory.Inventory) error
	Revert(inventory *inventory.Inventory) error
	Name() string
}
