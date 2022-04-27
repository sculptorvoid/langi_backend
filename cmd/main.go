package main

import (
	"context"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	langi "github.com/sculptorvoid/langi_backend"
	"github.com/sculptorvoid/langi_backend/pkg/handler"
	"github.com/sculptorvoid/langi_backend/pkg/repository"
	"github.com/sculptorvoid/langi_backend/pkg/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
	"os/signal"
	"syscall"
)

// @title                       Langi backend
// @version                     1.0
// @description                 Create your own unlimited dictionaries and learn foreign languages
// @license.name                MIT
// @license.url                 https://opensource.org/licenses/MIT
// @host                        localhost:8000
// @BasePath                    /
// @securityDefinitions.apikey  ApiKeyAuth
// @in                          header
// @name                        Authorization
func main() {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.PostgresConnect(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(langi.Server)

	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error starting http server: %s", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error server shutdown: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("error closing database: %s", err.Error())
	}
}
