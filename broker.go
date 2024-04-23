package main

import (
	"bytes"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

type Broker interface {
	Connect() error
	Close() error
	Subscribe(topic Topic, handler Handler) error
}

type RabbitMQBroker struct {
	dns         string
	ctx         context.Context
	timeout     int
	concurrency int
	connection  *amqp.Connection
	channel     *amqp.Channel
	start       time.Time
	Broker
}

func NewRabbitMQBroker(ctx context.Context, dns string, timeout, concurrency int) *RabbitMQBroker {
	return &RabbitMQBroker{
		dns:         dns,
		ctx:         ctx,
		timeout:     timeout,
		concurrency: concurrency,
	}
}

func (b *RabbitMQBroker) Connect() error {
	conn, err := amqp.Dial(b.dns)
	if err != nil {
		return err
	}

	b.connection = conn

	ch, err := conn.Channel()
	if err != nil {
		return err
	}

	b.channel = ch

	return nil
}

func (b *RabbitMQBroker) Close() error {
	if b.channel != nil {
		if err := b.channel.Close(); err != nil {
			return err
		}
	}

	if b.connection != nil {
		if err := b.connection.Close(); err != nil {
			return err
		}
	}

	return nil
}

func (b *RabbitMQBroker) Subscribe(topic Topic, handler Handler) error {
	ctx, cancel := context.WithTimeout(b.ctx, time.Duration(b.timeout)*time.Second)
	defer cancel()

	q, err := b.channel.QueueDeclare(string(topic), false, false, false, false, nil)
	if err != nil {
		return err
	}

	msgs, err := b.channel.Consume(q.Name, "", true, false, false, true, nil)
	if err != nil {
		return err
	}

	var tasks []Task
	for d := range msgs {
		tasks = append(tasks, NewConsumerTask(bytes.NewReader(d.Body), handler))
		d.Ack(false)
	}

	worker := NewDefaultWorker(tasks, 12)
	worker.Run()

	ctx.Done()

	return nil
}
