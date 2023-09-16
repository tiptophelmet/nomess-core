package pubsub

import (
	"sync"

	"github.com/tiptophelmet/nomess-core/v2/logger"
	"github.com/tiptophelmet/nomess-core/v2/pubsub/broker"
)

type pubSubClient struct {
	broker broker.PubSubBroker
	mu     sync.Mutex
}

var client *pubSubClient

func Init(driver, url string) {
	switch driver {
	case "redis":
		client = &pubSubClient{broker: broker.InitRedisBroker()}
	case "nats":
		client = &pubSubClient{broker: broker.InitNATSBroker()}
	default:
		logger.Panic("unsupported pubsub.driver: %v", driver)
	}

	client.broker.Connect(url)
}

func Connection() broker.PubSubBroker {
	client.mu.Lock()
	defer client.mu.Unlock()

	return client.broker
}
