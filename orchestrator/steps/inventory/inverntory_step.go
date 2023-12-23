package inventory

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/steps"
	"github.com/klauskie/saga-dt/orchestrator/steps/inventory/requesters"
)

type step struct {
	inventory  models.Inventory
	requester  requesters.Requester
	stepStatus steps.Status
}

func NewStep(task models.Task) steps.Step {
	return &step{
		inventory: models.Inventory{
			OrderId:    task.OrderId,
			UserId:     task.UserId,
			ProductId:  task.ProductId,
			OutOfStock: false,
		},
		stepStatus: steps.Pending,
		requester:  requesters.NewInventoryGrpcRequester("4042"),
	}
}

func (s step) Process() error {
	err := s.requester.Process(&s.inventory)
	if err != nil {
		s.stepStatus = steps.Failed
		return err
	}
	if s.inventory.OutOfStock {
		s.stepStatus = steps.Failed
		return fmt.Errorf("inventory step process failed. product unavailable")
	}

	s.stepStatus = steps.Complete
	return nil
}

func (s step) Revert() error {
	return s.requester.Revert(&s.inventory)
}

func (s step) Status() steps.Status {
	return s.stepStatus
}

func (s step) Name() string {
	return fmt.Sprintf("inventory/%s", s.requester.Name())
}
