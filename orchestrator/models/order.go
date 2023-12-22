package models

// Order orchestrator request
type Order struct {
	Id        string  `json:"id"`
	UserId    int     `json:"user_id"`
	ProductId int     `json:"product_id"`
	Amount    float64 `json:"amount"`
}
