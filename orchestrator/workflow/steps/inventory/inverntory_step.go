package inventory

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/inventory/requesters"
)

type step struct {
	inventory  Inventory
	requester  requesters.Requester
	stepStatus steps.Status
}

func NewStep(task models.Task) steps.Step {
	return &step{
		inventory: Inventory{
			OrderId:    task.OrderId,
			UserId:     task.UserId,
			ProductId:  task.ProductId,
			OutOfStock: false,
		},
		stepStatus: steps.Pending,
		requester:  requesters.NewInventoryGrpcRequester(),
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
