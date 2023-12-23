package workflow

import (
	"errors"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/steps"
	"github.com/klauskie/saga-dt/orchestrator/steps/inventory"
	"github.com/klauskie/saga-dt/orchestrator/steps/payments"
)

func Handle(task models.Task) error {
	stepList := workflowSteps(task)

	err := processSteps(stepList)
	if err != nil {
		return revertSteps(stepList)
	}
	return err
}

func processSteps(stepList []steps.Step) error {
	for _, step := range stepList {
		err := step.Process()
		if err != nil {
			return err
		}
	}

	return nil
}

func revertSteps(stepList []steps.Step) error {
	var errList error
	for _, step := range stepList {
		if step.Status() == steps.Complete {
			continue
		}
		if err := step.Revert(); err != nil {
			errList = errors.Join(errList, err)
		}
	}
	return errList
}

func workflowSteps(task models.Task) []steps.Step {
	return []steps.Step{
		inventory.NewStep(task),
		payments.NewStep(task),
	}
}
