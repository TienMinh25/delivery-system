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
	Publish(ctx context.Context, topic string, request []byte) error
}

type SubscriptionInfo struct {
	Topic    string
	GroupID  string // Consumer group
	Callback func(context.Context, Span, MessageQueue) error
}

type MessageQueue interface {
	Body() []byte
}
