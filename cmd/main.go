package main

import (
	"log"
	server "med"
	"med/configs"
	"med/pkg/handler"
	route "med/pkg/routes"
)

func main() {
	config := configs.InitConfig(*configs.NewConfigInfo())

	handlers := new(handler.Handler)
	routes := route.InitRoutes(handlers)

	server := new(server.Server)
	if err := server.Run(config.Server.Host, config.Server.Port, routes); err != nil {
		log.Fatal()
	}
}
