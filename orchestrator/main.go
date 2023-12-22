package main

import (
	"fmt"
	"github.com/klauskie/saga-dt/orchestrator/models"
	"github.com/klauskie/saga-dt/orchestrator/workflow"
)

func main() {
	fmt.Println("Orchestrator is up...")

	workflow.Handle(models.Order{
		Id:        "",
		UserId:    0,
		ProductId: 0,
		Amount:    0,
	})
}
