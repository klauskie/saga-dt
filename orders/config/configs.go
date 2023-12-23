package config

import "github.com/klauskie/saga-dt/orders/extensions/pubsub"

func DefaultEnv() Env {
	brokerHost := ""
	ps := pubsub.NewKafkaPub(brokerHost)

	return Env{
		Settings: struct{ BrokerHost string }{BrokerHost: brokerHost},
		X:        struct{ PubSub pubsub.IPubSub }{PubSub: ps},
	}
}
