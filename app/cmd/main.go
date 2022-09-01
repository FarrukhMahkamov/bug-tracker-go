package main

import (
	"log"

	server "github.com/FarrukhMahkamov/bugtracker"
	"github.com/FarrukhMahkamov/bugtracker/internal/handler"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
	"github.com/FarrukhMahkamov/bugtracker/internal/service"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	if err := InitConfig(); err != nil {
		log.Fatalf("Error while reading config: %s", err.Error())
	}
	logrus.SetFormatter(new(logrus.JSONFormatter))

	repositories := repository.NewRepository()
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	if err := srv.Serve(viper.GetString("port"), handlers.InitiRoutes()); err != nil {
		log.Fatalf("Error ocured while serivng server")
	}
}

func InitConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
