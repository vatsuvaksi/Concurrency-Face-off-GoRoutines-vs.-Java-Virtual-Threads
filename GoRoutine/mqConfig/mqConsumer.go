package mqconfig

import (
	"log"
	"sync"
	"time"

	"github.com/streadway/amqp"
)

type Consumer struct {
	rabbitMQ       *RabbitMQConfig
	messageHandler func([]amqp.Delivery)
}

// Init initializes the consumer and starts message consumption.
func (c *Consumer) Init(qos int) {
	log.Println("Initializing consumer...")

	messages, err := c.rabbitMQ.channel.Consume(
		c.rabbitMQ.queue, "", false, false, false, false, nil,
	)
	if err != nil {
		log.Fatalf("Failed to start consuming messages: %v", err)
	}

	c.rabbitMQ.StartDynamicQosAdjustment(1 * time.Minute)

	c.processMessages(messages, qos)
}

func (c *Consumer) processMessages(messages <-chan amqp.Delivery, qos int) {
	var wg sync.WaitGroup
	workerPool := make(chan struct{}, qos)
	buffer := make([]amqp.Delivery, 0, qos)

	for msg := range messages {
		buffer = append(buffer, msg)
		if len(buffer) >= qos {
			c.processBuffer(buffer, &wg, workerPool)
			buffer = make([]amqp.Delivery, 0, qos)
		}
	}

	if len(buffer) > 0 {
		c.processBuffer(buffer, &wg, workerPool)
	}

	wg.Wait()
}

func (c *Consumer) processBuffer(buffer []amqp.Delivery, wg *sync.WaitGroup, workerPool chan struct{}) {
	workerPool <- struct{}{}
	wg.Add(1)
	go func(buf []amqp.Delivery) {
		defer wg.Done()
		defer func() { <-workerPool }()
		c.messageHandler(buf)
	}(buffer)
}

// NewConsumer creates a new Consumer instance.
func NewConsumer(rabbitMQ *RabbitMQConfig, handler func([]amqp.Delivery)) *Consumer {
	return &Consumer{
		rabbitMQ:       rabbitMQ,
		messageHandler: handler,
	}
}
