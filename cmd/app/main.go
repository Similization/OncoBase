package main

import (
	"context"
	server "med"
	_ "med/docs"
	"med/pkg/config"
	"med/pkg/handler"
	"med/pkg/repository"
	route "med/pkg/routes"
	services "med/pkg/service"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

// @title OncomarkerAPI
// @version 1.0
// @description API for managing data on oncological markers.

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	logger := zerolog.New(os.Stdout).Level(zerolog.DebugLevel).With().Timestamp().Logger()

	config := config.InitConfig(*config.DefaultConfigInfo())

	db, err := repository.NewPostgresDB(&config.Database)
	if err != nil {
		logger.Error()
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
