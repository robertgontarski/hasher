package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"os"
	"time"
)

func main() {
	if err := godotenv.Load(); err != nil {
		slog.Error(fmt.Sprintf("error loading .env file: %v", err))
		return
	}

	setupLogger()

	ctx := context.Background()

	mysql := NewMysqlDatabase(DatabaseConf{
		Driver:       os.Getenv("DATABASE_DRIVER"),
		Dns:          os.Getenv("DATABASE_CONN"),
		MaxLifeTime:  time.Minute * 60,
		MaxIdleConns: 10,
		MaxOpenConns: 10,
	})

	if err := mysql.Connect(); err != nil {
		slog.Error("failed to connect to mysql: %v", err)
		return
	}

	defer func() {
		if err := mysql.Close(); err != nil {
			slog.Error("failed to close mysql: %v", err)
			return
		}
	}()

	store := NewMysqlStore(ctx, mysql)

	switch os.Getenv("APP_VERSION") {
	case RabbitMQAppVersion:
		NewRabbitMQVersion(ctx, store).InitVersion()
		return
	case HttpAppVersion:
		NewHttpVersion(ctx, store).InitVersion()
		return
	case GrpcAppVersion:
		NewGrpcVersion(ctx, store).InitVersion()
	default:
		slog.Error(fmt.Sprintf("app version not found, please check your APP_VERSION env variable, available options: (%s, %s, %s)", RabbitMQAppVersion, HttpAppVersion, GrpcAppVersion))
	}
}
