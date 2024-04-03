package main

import (
	"context"
	server "med"
	"med/pkg/config"
	"med/pkg/handler"
	"med/pkg/repository"
	route "med/pkg/routes"
	"med/pkg/services"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	config := config.InitConfig(*config.DefaultConfigInfo())

	db, err := repository.NewPostgresDB(&config.Database)
	if err != nil {
		logger.Error().Msg(err.Error())
	}

	repository := repository.NewRepository(db)
	service := services.NewService(*repository)
	handler := handler.NewHandler(service)

	routes := route.InitRoutes(handler)

	server := new(server.Server)
	go func() {
		if err := server.Run(config.Server.Host, config.Server.Port, routes); err != nil {
			logger.Error()
		}
	}()

	logger.Print("TodoApp Started")

	quitch := make(chan os.Signal, 1)
	signal.Notify(quitch, syscall.SIGTERM, syscall.SIGINT)
	<-quitch

	logger.Print("TodoApp Shutting Down")

	if err := server.Shutdown(context.Background()); err != nil {
		logger.Error().Msgf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logger.Error().Msgf("error occured on db connection close: %s", err.Error())
	}
}
