package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
)

type Requester interface {
	Process(inventory *models.Inventory) error
	Revert(inventory *models.Inventory) error
	Name() string
}
