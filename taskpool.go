package main

import "io"

type Task interface {
	Process() error
}

type ConsumerTask struct {
	Reader  io.Reader
	Handler Handler
}

func NewConsumerTask(r io.Reader, handler Handler) *ConsumerTask {
	return &ConsumerTask{
		Reader:  r,
		Handler: handler,
	}
}

func (t *ConsumerTask) Process() error {
	return t.Handler.Handle(t.Reader)
}

type SubscribeTask struct {
	Topic   Topic
	Broker  Broker
	Handler Handler
}

func NewSubscribeTask(topic Topic, broker Broker, handler Handler) *SubscribeTask {
	return &SubscribeTask{
		Topic:   topic,
		Broker:  broker,
		Handler: handler,
	}
}

func (t *SubscribeTask) Process() error {
	return t.Broker.Subscribe(t.Topic, t.Handler)
}
