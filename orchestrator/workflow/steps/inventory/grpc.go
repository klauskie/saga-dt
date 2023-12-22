package inventory

import "github.com/klauskie/saga-dt/orchestrator/workflow/steps"

type inventoryGrpcStep struct{}

func (s inventoryGrpcStep) Process() bool { return false }
func (s inventoryGrpcStep) Revert() bool  { return false }
func (s inventoryGrpcStep) Status() steps.Status {
	return steps.Complete
}
func (s inventoryGrpcStep) Name() string {
	return "inventory/grpc"
}
