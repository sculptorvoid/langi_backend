package main

import (
	langi "github.com/sculptorvoid/langi_backend"
	"github.com/sculptorvoid/langi_backend/pkg/handler"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
	"github.com/sculptorvoid/langi_backend/pkg/service"
	"log"
)

func main() {

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(langi.Server)
	if err := server.Run("8000", handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting http server: %s", err.Error())
	}
}
