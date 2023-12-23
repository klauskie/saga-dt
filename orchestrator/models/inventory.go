package models

type Inventory struct {
	OrderId    string
	UserId     int
	ProductId  int
	OutOfStock bool
}
