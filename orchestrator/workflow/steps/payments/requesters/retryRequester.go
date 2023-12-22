package requesters

import (
	"github.com/klauskie/saga-dt/orchestrator/workflow/steps/payments"
	"time"
)

type RetryRequester struct {
	DefaultRequester Requester
	MaxRetries       int
	Delay            time.Duration
}

func (rr RetryRequester) Process(payment *payments.Payment) (err error) {
	for i := 0; i < rr.MaxRetries; i++ {
		err = rr.DefaultRequester.Process(payment)
		if err == nil {
			return err
		}
		<-time.After(rr.Delay)
	}
	return err
}

func (rr RetryRequester) Revert(payment *payments.Payment) (err error) {
	for i := 0; i < rr.MaxRetries; i++ {
		err = rr.DefaultRequester.Revert(payment)
		if err == nil {
			return err
		}
		<-time.After(rr.Delay)
	}
	return err
}
