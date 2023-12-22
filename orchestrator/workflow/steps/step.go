package steps

type Step interface {
	Name() string
	Process() bool
	Revert()
	Status() Status
}

type Status string

const (
	Pending  Status = "pending"
	Complete Status = "complete"
	Failed   Status = "failed"
)