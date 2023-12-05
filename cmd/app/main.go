package main

import (
	server "med"
	"med/pkg/config"
	"med/pkg/handler"
	"med/pkg/repository"
	route "med/pkg/routes"
	"med/pkg/services"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
)

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
	if err := server.Run(config.Server.Host, config.Server.Port, routes); err != nil {
		logger.Error()
	}
}
