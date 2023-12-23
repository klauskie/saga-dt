package config

import "github.com/klauskie/saga-dt/orders/extensions/pubsub"

type Env struct {
	Settings struct {
		BrokerHost string
	}
	X struct {
		PubSub pubsub.IPubSub
	}
}
