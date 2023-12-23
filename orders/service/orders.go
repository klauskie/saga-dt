package service

import (
	"bytes"
	"context"
	"encoding/json"
	"github.com/klauskie/saga-dt/orders/config"
	"github.com/klauskie/saga-dt/orders/extensions/pubsub"
	"github.com/klauskie/saga-dt/orders/models"
)

type IOrderService interface {
	Submit(order models.Order) error
}

type orderService struct {
	pubSub pubsub.IPubSub
}

func NewOrderService(env config.Env) IOrderService {
	return &orderService{
		pubSub: env.X.PubSub,
	}
}

func (o *orderService) Submit(order models.Order) error {
	b, err := json.Marshal(order)
	if err != nil {
		return err
	}

	return o.pubSub.Produce(context.Background(), "orders", bytes.NewReader(b))
}
