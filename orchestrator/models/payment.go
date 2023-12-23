package models

type Payment struct {
	UserId   int
	OrderId  string
	Amount   float64
	Currency string
	Status   PaymentStatus
}

type PaymentStatus string

const (
	PaymentNone     PaymentStatus = "none"
	PaymentApproved PaymentStatus = "approved"
	PaymentRejected PaymentStatus = "rejected"
)
