package main

import (
	"fmt"
	"log"
	"os"
	"time"

	todo "github.com/Glebegor/do-app"
	handler "github.com/Glebegor/do-app/pkg/handler"
	repository "github.com/Glebegor/do-app/pkg/repository"
	service "github.com/Glebegor/do-app/pkg/service"
	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"
	viper "github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("Error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	db, err := repository.NewPostgresDb(repository.Config{
		Host:     viper.GetString("db.Host"),
		Port:     viper.GetString("db.Port"),
		Username: viper.GetString("db.Username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.DBName"),
		SSLMode:  viper.GetString("db.SSLMode"),
	})
	if err != nil {
		log.Fatalf("failed to initializate db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	fmt.Printf("%s: Server is running on port 8000. BTW", time.Now())
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
