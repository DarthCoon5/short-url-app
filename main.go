package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"math/rand"
	"short-url-app/pkg/handler"
	"short-url-app/pkg/repository"
	"short-url-app/pkg/service"
	"short-url-app/run"
	"time"
)

func main() {
	if err := initConfig(); err != nil {
		logrus.Fatalf("Init configs error: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DbName:   viper.GetString("db.dbname"),
		SslMode:  viper.GetString("db.sslmode"),
		Password: viper.GetString("db.password"),
	})
	if err != nil {
		logrus.Fatalf("Init db error: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos, fmt.Sprintf("http://localhost:%s/", viper.GetString("port")))
	handlers := handler.NewHandler(services)

	rand.Seed(time.Now().UnixNano())

	srv := new(run.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Run server error: %s", err.Error())
	}

}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
