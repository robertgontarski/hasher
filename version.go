package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
)

type Version interface {
	InitVersion()
}

type RabbitMQVersion struct {
	Ctx   context.Context
	Store Store
	Version
}

func NewRabbitMQVersion(ctx context.Context, store Store) *RabbitMQVersion {
	return &RabbitMQVersion{
		Ctx:   ctx,
		Store: store,
	}
}

func (v *RabbitMQVersion) InitVersion() {
	timeout, err := strconv.Atoi(os.Getenv("RABBITMQ_CONN_TIMEOUT"))
	if err != nil {
		slog.Error(fmt.Sprintf("error while converting to int: %v", err))
	}

	concurrency, err := strconv.Atoi(os.Getenv("RABBITMQ_CONN_CONCURRENCY"))
	if err != nil {
		slog.Error(fmt.Sprintf("error while converting to int: %v", err))
	}

	rabbit := NewRabbitMQBroker(v.Ctx, os.Getenv("RABBITMQ_CONN"), timeout, concurrency)
	if err := rabbit.Connect(); err != nil {
		slog.Error(fmt.Sprintf("failed to connect to rabbitmq: %v", err))
		return
	}

	defer func() {
		if err := rabbit.Close(); err != nil {
			slog.Error("failed to close rabbitmq: %v", err)
			return
		}
	}()

	tasks := []Task{
		NewSubscribeTask(EmailChangeQueue, rabbit, NewEmailHandler(v.Store, NewEmailHasher())),
		NewSubscribeTask(PhoneChangeQueue, rabbit, NewPhoneHandler(v.Store, NewPhoneHasher())),
		NewSubscribeTask(NameChangeQueue, rabbit, NewNameHandler(v.Store, NewNameHasher())),
	}

	wp := NewDefaultWorker(tasks, 3)
	wp.Run()
}

type HttpVersion struct {
	Ctx   context.Context
	Store Store
	Version
}

func NewHttpVersion(ctx context.Context, store Store) *HttpVersion {
	return &HttpVersion{
		Ctx:   ctx,
		Store: store,
	}
}

func (v *HttpVersion) InitVersion() {
	srv := NewHttpServer(v.Ctx, os.Getenv("HTTP_SERVER_ADDR"), v.Store)
	if err := srv.Listen(); err != nil {
		slog.Error(fmt.Sprintf("can not start server: %v", err))
	}
}

type GrpcVersion struct {
	Ctx   context.Context
	Store Store
	Version
}

func NewGrpcVersion(ctx context.Context, store Store) *GrpcVersion {
	return &GrpcVersion{
		Ctx:   ctx,
		Store: store,
	}
}

func (v *GrpcVersion) InitVersion() {
	srv := NewGrpcServer(v.Ctx, os.Getenv("GRPC_SERVER_ADDR"), v.Store)
	if err := srv.Listen(); err != nil {
		slog.Error(fmt.Sprintf("can not start server: %v", err))
	}
}
