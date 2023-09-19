package broker

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/tiptophelmet/nomess-core/v5/errs"
	"github.com/tiptophelmet/nomess-core/v5/logger"
)

func InitRedisBroker() *RedisBroker {
	return &RedisBroker{subscriptions: make(map[string]*redis.PubSub)}
}

type RedisBroker struct {
	client        *redis.Client
	subscriptions map[string]*redis.PubSub
}

func (r *RedisBroker) Connect(url string) error {
	options, err := redis.ParseURL(url)
	if err != nil {
		logger.Error("failed to connect to redis (pubsub): %v", err.Error())
		return errs.ErrPubSubBrokerConnectionFailed
	}

	r.client = redis.NewClient(options)
	return nil
}

func (r *RedisBroker) Publish(topic string, message []byte) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	intCmd := r.client.Publish(ctx, topic, message)
	return intCmd.Err()
}

func (r *RedisBroker) Subscribe(topic string, handler func(message []byte)) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	sub := r.client.Subscribe(ctx, topic)

	_, err := sub.Receive(ctx)
	if err != nil {
		return err
	}

	r.subscriptions[topic] = sub

	ch := sub.Channel()
	for msg := range ch {
		handler([]byte(msg.Payload))
	}

	return nil
}

func (r *RedisBroker) Unsubscribe(topic string) error {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	sub, found := r.subscriptions[topic]
	if found {
		delete(r.subscriptions, topic)
		return sub.Unsubscribe(ctx)
	}

	return nil
}

func (r *RedisBroker) IsConnected() bool {
	ctx, ctxCancel := context.WithCancel(context.Background())
	defer ctxCancel()

	statusCmd := r.client.Ping(ctx)
	if err := statusCmd.Err(); err != nil {
		return false
	}

	return true
}

func (r *RedisBroker) Close() error {
	return r.client.Close()
}
