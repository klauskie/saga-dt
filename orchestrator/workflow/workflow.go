package workflow

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps"
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"
)

func Handle(r models.Order) {
	stepList := workflowSteps(r)

	err := processSteps(stepList)
	if err != nil {
		revertSteps(stepList)
	}
}

func processSteps(stepList []steps.Step) error {
	for _, step := range stepList {
		ok := step.Process() // TODO add a retry
		if !ok {
			return fmt.Errorf("step %s failed", step.Name())
		}
	}

	return nil
}

func revertSteps(stepList []steps.Step) {
	for _, step := range stepList {
		if step.Status() == steps.Complete {
			continue
		}
		step.Revert() // TODO add a retry
	}
}

func workflowSteps(r models.Order) []steps.Step {
	return []steps.Step{payments.NewGrpcStep(r)}
}
