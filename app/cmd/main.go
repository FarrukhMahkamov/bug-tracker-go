package main

import (
	server "github.com/FarrukhMahkamov/bugtracker"
	"github.com/FarrukhMahkamov/bugtracker/internal/handler"
	"github.com/FarrukhMahkamov/bugtracker/internal/repository"
	"github.com/FarrukhMahkamov/bugtracker/internal/service"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	

	if err := InitConfig(); err != nil {
		logrus.Fatalf("Error while reading config: %s", err.Error())
	}
	logrus.SetFormatter(new(logrus.JSONFormatter))

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: viper.GetString("db.password"),
		DBname:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})

	if err != nil {
		logrus.Fatalf("Error while initializing database: %s", err.Error())
	}

	repositories := repository.NewRepository(db)
	services := service.NewService(repositories)
	handlers := handler.NewHandler(services)
	srv := new(server.Server)

	if err := srv.Serve(viper.GetString("port"), handlers.InitiRoutes()); err != nil {
		logrus.Fatalf("Error ocured while serivng server")
	}
}

func InitConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
