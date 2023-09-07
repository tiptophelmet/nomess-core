package broker

type PubSubBroker interface {
	Connect(url string) error

	Publish(topic string, message []byte) error
	Subscribe(topic string, handler func(message []byte)) error
	Unsubscribe(topic string) error

	IsConnected() bool
	Close() error
}
