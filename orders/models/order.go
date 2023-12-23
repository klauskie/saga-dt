package models

type Order struct {
	OrderId  string `json:"order_id"`
	UserId   int    `json:"user_id"`
	Products []struct {
		ProductId int     `json:"product_id"`
		Price     float64 `json:"price"`
	} `json:"products"`
}
