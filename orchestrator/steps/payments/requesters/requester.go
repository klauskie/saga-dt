package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
)

type Requester interface {
	Process(payment *models.Payment) error
	Revert(payment *models.Payment) error
	Name() string
}
