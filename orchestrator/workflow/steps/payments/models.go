package payments

type Payment struct {
	UserId   int
	OrderId  string
	Amount   float64
	Currency string
	Status   PaymentStatus
}

type PaymentStatus string

const (
	None     PaymentStatus = "none"
	Approved PaymentStatus = "approved"
	Rejected PaymentStatus = "rejected"
)
