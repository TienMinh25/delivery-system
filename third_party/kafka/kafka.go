package kafka

import (
	"context"
	"fmt"
	"github.com/TienMinh25/delivery-system/internal/env"
	"github.com/TienMinh25/delivery-system/pkg"
	"github.com/TienMinh25/delivery-system/pkg/worker"
	kafkaconfluent "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	"sync"
	"time"
)

type queue struct {
	serviceName string                             // used for distributed tracing
	groupID     string                             // group consumer id
	producer    *kafkaconfluent.Producer           // producer
	consumer    *kafkaconfluent.Consumer           // consumer
	subscribers map[string][]*pkg.SubscriptionInfo // information of subscribers
	mu          sync.RWMutex                       // concurrent lock when subscribe
	workerPool  *worker.Pool                       // worker pool to process message
	tracer      pkg.DistributedTracer              // tracer
	ctx         context.Context
	cancel      context.CancelFunc
	wg          sync.WaitGroup
}

func NewQueue(serviceName string, tracer pkg.DistributedTracer, config *env.EnvManager) (pkg.Queue, error) {
	ctx, cancel := context.WithCancel(context.Background())

	q := &queue{
		groupID:     config.Kafka.KafkaGroupID,
		ctx:         ctx,
		cancel:      cancel,
		subscribers: make(map[string][]*pkg.SubscriptionInfo, 0),
		tracer:      tracer,
	}

	producer, err := newKafkaProducer(config.Kafka.KafkaBrokers, config.Kafka.KafkaRetryAttempts, config.Kafka.KafkaProducerMaxWait)

	if err != nil {
		cancel()
		return nil, fmt.Errorf("failed to create producer: %w", err)
	}
	q.producer = producer

	q.consumer, err = newKafkaConsumer(config.Kafka.KafkaBrokers, fmt.Sprintf("%s-%s", config.Kafka.KafkaGroupID, serviceName), config.Kafka.KafkaConsumerFetchMinBytes, config.Kafka.KafkaConsumerFetchMaxBytes, config.Kafka.KafkaConsumerMaxWait)
	if err != nil {
		q.producer.Close()
		cancel()
		return nil, fmt.Errorf("failed to create consumer: %w", err)
	}

	q.workerPool = worker.NewPool(config.ServiceWorkerPool.CapacityWorkerPool, config.ServiceWorkerPool.MessageSize, q.processMessage)

	// Start the message consumer loop
	// consume all message from all topic which is subscribed
	q.wg.Add(1)
	go q.consumeMessages()

	return q, nil
}

func newKafkaProducer(brokers string, retries, producerMaxWait int) (*kafkaconfluent.Producer, error) {
	p, err := kafkaconfluent.NewProducer(&kafkaconfluent.ConfigMap{
		"bootstrap.servers":                     brokers,
		"client.id":                             uuid.New().String(),
		"acks":                                  "all",
		"enable.idempotence":                    true,
		"max.in.flight.requests.per.connection": 5,
		"retries":                               retries,
		"linger.ms":                             producerMaxWait,
		"transactional.id":                      uuid.New().String(),
	})

	if err != nil {
		return nil, err
	}

	if err = p.InitTransactions(context.Background()); err != nil {
		p.Close()
		return nil, fmt.Errorf("failed to initialize transactions: %w", err)
	}

	go func() {
		for e := range p.Events() {
			switch ev := e.(type) {
			case *kafkaconfluent.Message:
				if ev.TopicPartition.Error != nil {
					fmt.Printf("Failed to deliver message: %v\n", ev.TopicPartition)
				} else {
					fmt.Printf("Successfully produced record to topic %s partition [%d] @ offset %v\n",
						*ev.TopicPartition.Topic, ev.TopicPartition.Partition, ev.TopicPartition.Offset)
				}
			}
		}
	}()

	return p, nil
}

func newKafkaConsumer(brokers, groupID string, fetchMinBytes, fetchMaxBytes, timeMaxWait int) (*kafkaconfluent.Consumer, error) {
	c, err := kafkaconfluent.NewConsumer(&kafkaconfluent.ConfigMap{
		"bootstrap.servers":  brokers,
		"client.id":          uuid.New().String(),
		"group.id":           groupID,
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
		// Consumer Tuning
		"max.poll.interval.ms":  60000, // 1p (kafka will kick consumer)
		"heartbeat.interval.ms": 5000,  // 5s
		"session.timeout.ms":    45000,
		"fetch.min.bytes":       fetchMinBytes,
		"fetch.max.bytes":       fetchMaxBytes,
		"fetch.wait.max.ms":     timeMaxWait,
		"isolation.level":       "read_committed",
	})

	if err != nil {
		return nil, err
	}

	return c, nil
}

// Close implements pkg.Queue.
func (q *queue) Close() error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.cancel() // Hủy tất cả goroutine (kể cả đang consume hay reconnect)
	q.workerPool.GracefulShutdown()

	// Wait for all goroutines to finish
	q.wg.Wait()

	// Close Kafka connections
	if err := q.consumer.Close(); err != nil {
		return err
	}
	q.producer.Close()
	return nil
}

// Publish implements pkg.Queue.
func (q *queue) Produce(ctx context.Context, topic string, payload []byte) error {
	// Start a span for tracing if tracer is provided
	var span pkg.Span

	if q.tracer != nil {
		ctx, span = q.tracer.StartFromContext(ctx, "kafka-produce")
		defer span.End()
	}

	// begin transaction
	err := q.producer.BeginTransaction()

	if err != nil {
		fmt.Printf("Failed to begin transaction: %s\n", err)
		return err
	}

	// Prepare message
	message := &kafkaconfluent.Message{
		TopicPartition: kafkaconfluent.TopicPartition{
			Topic:     &topic,
			Partition: kafkaconfluent.PartitionAny,
		},
		Value: payload,
	}

	// Inject tracing headers if span is available
	if span != nil {
		q.tracer.Inject(ctx, &KafkaHeadersCarrier{
			headers: message.Headers,
		})
	}

	if err = q.producer.Produce(message, nil); err != nil {
		fmt.Printf("Failed to produce message: %s\n", err)
		_ = q.producer.AbortTransaction(context.Background())
		return err
	}

	select {
	case <-ctx.Done():
		_ = q.producer.AbortTransaction(ctx)
		return ctx.Err()
	}

	// Commit transaction
	if err := q.producer.CommitTransaction(ctx); err != nil {
		_ = q.producer.AbortTransaction(ctx)
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Subscribe implements pkg.Queue.
// chi subscribe topic cho consumer -> chi ra rang group consumer do se poll luon message tu topic do tren kafka ve
func (q *queue) Subscribe(ctx context.Context, payload *pkg.SubscriptionInfo) error {
	if payload == nil || payload.Topic == "" || payload.Callback == nil {
		return fmt.Errorf("invalid subscription info")
	}

	q.mu.Lock()
	defer q.mu.Unlock()

	// Store subscription info
	q.subscribers[payload.Topic] = append(q.subscribers[payload.Topic], payload)

	topics := make([]string, 0, len(q.subscribers))

	for topic := range q.subscribers {
		topics = append(topics, topic)
	}

	return q.consumer.SubscribeTopics(topics, nil)
}

// consumeMessages is the main consumer loop
func (q *queue) consumeMessages() {
	defer q.wg.Done()

	for {
		select {
		case <-q.ctx.Done():
			// Context was cancelled, exit the loop
			return
		default:
			// Poll for messages with a timeout
			ev := q.consumer.Poll(100)

			// no new message in kafka
			if ev == nil {
				continue
			}

			switch e := ev.(type) {
			case *kafkaconfluent.Message:
				// Push message to worker pool for processing
				q.workerPool.PushMessage(e)
			case kafkaconfluent.Error:
				fmt.Printf("Consumer error: %v\n", e)

				if e.Code() == kafkaconfluent.ErrAllBrokersDown {
					time.Sleep(5 * time.Second)
				}
			}
		}
	}
}

// processMessage processes a message received from Kafka
// check topic belong to the handlers and process that with handler
func (q *queue) processMessage(message interface{}) error {
	msg, ok := message.(*kafkaconfluent.Message)

	if !ok {
		return fmt.Errorf("invalid message type")
	}

	topic := *msg.TopicPartition.Topic

	// create a context with timeout for message processing
	ctx, cancel := context.WithTimeout(q.ctx, time.Second*30)
	defer cancel()

	// extract span
	var span pkg.Span

	if q.tracer != nil {
		// Create a carrier from Kafka headers
		carrier := NewKafkaHeadersCarrier(msg.Headers)

		// Extract span context from carrier
		extractedSpan := q.tracer.Extract(ctx, carrier)

		// Start a span from extracted context
		ctx, span = q.tracer.StartFromSpan(ctx, extractedSpan, q.serviceName)
		defer span.End()

		// Record topic as an attribute
		span.SetAttributes("kafka.topic", topic)
	}

	// find handlers for this topic
	q.mu.RLock()
	handlers := q.subscribers[topic]
	q.mu.RUnlock()

	if len(handlers) == 0 {
		// No handlers for this topic, commit message and return
		_, err := q.consumer.CommitMessage(msg)
		return err
	}

	// Process message with all registered handlers
	var lastErr error

	for _, handler := range handlers {
		if err := handler.Callback(ctx, span, msg); err != nil {
			lastErr = err

			// Record error in span
			if span != nil {
				span.RecordError(err)
			}

			// Log the error but continue processing with other handlers
			fmt.Printf("Error processing message for topic %s: %v\n", topic, err)
		}
	}

	// Commit the message offset
	_, err := q.consumer.CommitMessage(msg)

	if err != nil {
		if span != nil {
			span.RecordError(err)
		}
		return fmt.Errorf("failed to commit message: %w", err)
	}

	return lastErr
}

// KafkaHeadersCarrier implements TextMapCarrier for Kafka headers
type KafkaHeadersCarrier struct {
	headers []kafkaconfluent.Header
}

func NewKafkaHeadersCarrier(headers []kafkaconfluent.Header) *KafkaHeadersCarrier {
	return &KafkaHeadersCarrier{headers: headers}
}

func (c *KafkaHeadersCarrier) Get(key string) string {
	for _, h := range c.headers {
		if h.Key == key {
			return string(h.Value)
		}
	}
	return ""
}

func (c *KafkaHeadersCarrier) Set(key string, value string) {
	for i, h := range c.headers {
		if h.Key == key {
			c.headers[i].Value = []byte(value)
			return
		}
	}
	c.headers = append(c.headers, kafkaconfluent.Header{
		Key:   key,
		Value: []byte(value),
	})
}

func (c *KafkaHeadersCarrier) Keys() []string {
	keys := make([]string, len(c.headers))
	for i, h := range c.headers {
		keys[i] = h.Key
	}
	return keys
}
