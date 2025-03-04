package kafka

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/TienMinh25/delivery-system/pkg"
	kafkaconfluent "github.com/confluentinc/confluent-kafka-go/kafka"
)

// TODO: handle later
type queue struct {
	topic          string                   // topic
	groupID        string                   // group consumer id
	producer       *kafkaconfluent.Producer // producer
	consumer       *kafkaconfluent.Consumer // consumer
	subcribers     []*pkg.SubscriptionInfo  // information of subcriber
	ctx            context.Context
	cancel         context.CancelFunc    // Dùng để cancel tất cả goroutine cũ khi consumer reconnect
	reconnectCh    chan bool             // signal for reconnecting consumer
	tracer         pkg.DistributedTracer // distributed trace log
	reconnectDelay time.Duration         // reconnect delay time
	mu             sync.RWMutex          // concurrent lock when subcribe
}

func NewQueue(serviceName string, tracer pkg.DistributedTracer) (pkg.Queue, error) {
	retryDelayMs, err := strconv.Atoi(os.Getenv("KAFKA_RETRY_DELAY"))
	brokersString := os.Getenv("KAFKA_BROKERS")
	groupID := os.Getenv("KAFKA_GROUP_ID")

	if err != nil {
		// fallback value
		retryDelayMs = 2000
	}

	ctx, cancel := context.WithCancel(context.Background())

	q := &queue{
		topic:          os.Getenv("KAFKA_TOPIC"),
		groupID:        groupID,
		ctx:            ctx,
		cancel:         cancel,
		reconnectCh:    make(chan bool),
		producer:       newKafkaProducer(brokersString),
		consumer:       newKafkaConsumer(brokersString, groupID),
		subcribers:     make([]*pkg.SubscriptionInfo, 0),
		reconnectDelay: time.Duration(retryDelayMs) * time.Millisecond,
		tracer:         tracer,
	}

	// handle for case consumer died or disconnect
	go q.reconnectConsumer(brokersString, groupID)

	return q, nil
}

func newKafkaProducer(brokers string) *kafkaconfluent.Producer {
	retries, err := strconv.Atoi(os.Getenv("KAFKA_RETRY_ATTEMPTS"))

	if err != nil {
		// fallback value
		retries = 5
	}

	producerMaxWait, err := strconv.Atoi(os.Getenv("KAFKA_PRODUCER_MAX_WAIT"))

	if err != nil {
		// fallback value
		producerMaxWait = 300
	}

	p, err := kafkaconfluent.NewProducer(&kafkaconfluent.ConfigMap{
		"bootstrap.servers":                     brokers,
		"client.id":                             "myProducer",
		"acks":                                  "all",
		"enable.idempotence":                    true,
		"max.in.flight.requests.per.connection": 5,
		"retries":                               retries,
		"linger.ms":                             producerMaxWait,
	})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}

	return p
}

func newKafkaConsumer(brokers, groupID string) *kafkaconfluent.Consumer {
	fetchMinBytes, err := strconv.Atoi(os.Getenv("KAFKA_CONSUMER_FETCH_MIN_BYTES"))

	if err != nil {
		// fallback value
		fetchMinBytes = 5
	}

	fetchMaxBytes, err := strconv.Atoi(os.Getenv("KAFKA_CONSUMER_FETCH_MAX_BYTES"))

	if err != nil {
		// fallback value
		fetchMaxBytes = 1e6
	}

	timeMaxWait, err := strconv.Atoi(os.Getenv("KAFKA_CONSUMER_MAX_WAIT"))

	if err != nil {
		// fallback value
		timeMaxWait = 1000
	}

	c, err := kafkaconfluent.NewConsumer(&kafkaconfluent.ConfigMap{
		"bootstrap.servers":  brokers,
		"client.id":          "my-consumer",
		"group.id":           groupID,
		"enable.auto.commit": false,
		"auto.offset.reset":  "earliest",
		"fetch.min.bytes":    fetchMinBytes,
		"fetch.max.bytes":    fetchMaxBytes,
		"fetch.max.wait.ms":  timeMaxWait,
	})

	if err != nil {
		fmt.Printf("Failed to create consumer: %s\n", err)
		os.Exit(1)
	}

	return c
}

// Close implements pkg.Queue.
func (q *queue) Close() error {
	q.cancel() // Hủy tất cả goroutine (kể cả đang consume hay reconnect)
	q.mu.Lock()
	defer q.mu.Unlock()
	q.consumer.Close()
	q.producer.Close()
	return nil
}

// Publish implements pkg.Queue.
func (q *queue) Publish(ctx context.Context, topic string, request []byte) error {
	panic("unimplemented")
}

// Subscribe implements pkg.Queue.
func (q *queue) Subscribe(ctx context.Context, payload *pkg.SubscriptionInfo) error {
	q.mu.Lock()
	defer q.mu.Unlock()
	q.subcribers = append(q.subcribers, payload)
	go q.consume(payload)
	return nil
}

func (q *queue) reconnectConsumer(brokers, groupID string) {
	for {
		select {
		case <-q.ctx.Done():
			return
		case <-q.reconnectCh:
			q.mu.Lock()
			q.consumer.Close()
			q.consumer = newKafkaConsumer(brokers, groupID)
			q.mu.Unlock()

			for _, sub := range q.subcribers {
				go q.consume(sub)
			}
		}
	}
}

func (q *queue) consume(sub *pkg.SubscriptionInfo) {
	for {
		select {
		case <-q.ctx.Done():
			return
		default:
			msg, err := q.consumer.ReadMessage()
		}
	}
}

type messageQueue struct {
	body []byte
}

// Body implements pkg.MessageQueue.
func (m *messageQueue) Body() []byte {
	return m.body
}

func newMessageQueue() pkg.MessageQueue {
	return &messageQueue{}
}
