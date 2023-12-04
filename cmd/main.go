package main

import (
	"fmt"
	"log"
	server "med"
	"med/pkg/handler"
	route "med/pkg/routes"
)

func main() {
	handlers := new(handler.Handler)
	routes := route.InitRoutes(handlers)

	server := new(server.Server)
	fmt.Println("Server is running ...")
	if err := server.Run("8080", routes); err != nil {
		log.Fatal()
	}
}
