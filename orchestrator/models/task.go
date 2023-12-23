package models

// Task orchestrator request
type Task struct {
	OrderId   string  `json:"order_id"`
	UserId    int     `json:"user_id"`
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
}
