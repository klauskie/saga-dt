package pubsub

import (
	"context"
	"io"
)

type IPubSub interface {
	Produce(ctx context.Context, topic string, data io.Reader) error
}
