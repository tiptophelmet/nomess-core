package broker

import (
	"fmt"

	"github.com/nats-io/nats.go"
	"github.com/tiptophelmet/nomess-core/errs"
	"github.com/tiptophelmet/nomess-core/logger"
)

func InitNATSBroker() *NATSBroker {
	return &NATSBroker{subscriptions: make(map[string]*nats.Subscription)}
}

type NATSBroker struct {
	client        *nats.Conn
	subscriptions map[string]*nats.Subscription
}

func (n *NATSBroker) Connect(url string) error {
	var err error
	n.client, err = nats.Connect(url)

	if err != nil {
		logger.Error("failed to connect to NATS: %v", err.Error())
		return errs.ErrPubSubBrokerConnectionFailed
	}

	return nil
}

func (n *NATSBroker) Publish(topic string, message []byte) error {
	return n.client.Publish(topic, message)
}

func (n *NATSBroker) Subscribe(topic string, handler func(message []byte)) error {
	sub, err := n.client.Subscribe(topic, func(m *nats.Msg) {
		handler(m.Data)
	})

	if err != nil {
		return err
	}

	n.subscriptions[topic] = sub
	return nil
}

func (n *NATSBroker) Unsubscribe(topic string) error {
	sub, found := n.subscriptions[topic]
	if found {
		delete(n.subscriptions, topic)
		return sub.Unsubscribe()
	}

	return fmt.Errorf("failed to unsubscribe from NATS topic: %v", topic)
}

func (n *NATSBroker) IsConnected() bool {
	return n.client.IsConnected()
}

func (n *NATSBroker) Close() error {
	n.client.Close()
	if n.client.IsClosed() {
		return nil
	}

	return fmt.Errorf(errs.ErrPubSubBrokerConnectionNotClosed.Error()+" %v", "NATS")
}
