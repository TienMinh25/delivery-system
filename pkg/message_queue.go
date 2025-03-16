package pkg

import "context"

type Queue interface {
	Subscriber
	Publisher

	// Closes the connection to the message broker, returning an error if the operation fails.
	Close() error
}

type Subscriber interface {
	// Subcribes to a topic or queue for receiving messages, returning an error if the subcription fails.
	Subscribe(ctx context.Context, payload *SubscriptionInfo) error
}

type Publisher interface {
	// Publishes a message to a specified topic, returning an error if the operation fails.
	Produce(ctx context.Context, topic string, request []byte) error
}

type HandlerFunc func(ctx context.Context, span Span, message interface{}) error

type SubscriptionInfo struct {
	Topic    string
	Callback HandlerFunc
}
