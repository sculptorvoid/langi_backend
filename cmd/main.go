package main

import (
	langi "github.com/sculptorvoid/langi_backend"
	"github.com/sculptorvoid/langi_backend/pkg/handler"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
	"github.com/sculptorvoid/langi_backend/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(langi.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error starting http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
