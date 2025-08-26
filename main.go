package main

import (
	"context"
	"github.com/cynx-io/athena-ideas/internal/app"
	"github.com/cynx-io/athena-ideas/internal/dependencies/config"
	"github.com/cynx-io/athena-ideas/internal/grpc"
	"github.com/cynx-io/cynx-core/src/logger"
	"github.com/sirupsen/logrus"
)

func main() {

	ctx := context.Background()
	defer ctx.Done()

	config.Init()
	logLevel, err := logrus.ParseLevel(config.Config.Elastic.Level)
	if err != nil {
		logLevel = logrus.DebugLevel
	}
	logger.Init(logger.LoggerConfig{
		Level:            logLevel,
		ElasticsearchURL: []string{config.Config.Elastic.Url},
		ServiceName:      "athena-ideas",
	})

	logger.Info(ctx, "Starting athena")
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		cancel()
	}()

	logger.Info(ctx, "Initializing App")
	application, err := app.NewApp(ctx)
	if err != nil {
		panic(err)
	}

	logger.Info(ctx, "Initializing GRPC")
	grpcServer := grpc.Server{
		Services: application.Services,
	}

	err = grpcServer.Start(ctx)
	if err != nil {
		panic("failed to start grpc server " + err.Error())
	}
}
