package main

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow"
)

func main() {
	fmt.Println("Orchestrator is up...")

	workflow.Handle(models.Task{
		OrderId:   "123",
		UserId:    987,
		ProductId: 2,
		Amount:    20.99,
	})
}
