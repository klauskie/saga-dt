package main

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow"
)

func main() {
	fmt.Println("Orchestrator is up...")

	workflow.Handle(models.Task{
		OrderId:   "",
		UserId:    0,
		ProductId: 0,
		Amount:    0,
	})
}
