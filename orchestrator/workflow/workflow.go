package workflow

import (
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/inventory"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"
)

func Handle(task models.Task) {
	stepList := workflowSteps(task)

	err := processSteps(stepList)
	if err != nil {
		revertSteps(stepList)
	}
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

func revertSteps(stepList []steps.Step) {
	for _, step := range stepList {
		if step.Status() == steps.Complete {
			continue
		}
		_ = step.Revert()
	}
}

func workflowSteps(task models.Task) []steps.Step {
	return []steps.Step{
		inventory.NewStep(task),
		payments.NewStep(task),
	}
}
