package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

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

	go func() {
		if err := srv.Serve(viper.GetString("port"), handlers.InitiRoutes()); err != nil {
			logrus.Fatalf("Error ocured while serivng server")
		}
	}()

	logrus.Print("Bug Tracker shuting down")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Error occured while shuttingdown server : %s", err.Error())
	}

	if err := db.Close(); err != nil {
		logrus.Errorf("Error occured while closing db : %s", err.Error())
	}
}

func InitConfig() error {
	viper.AddConfigPath("./config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
