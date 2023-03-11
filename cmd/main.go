package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	todo "github.com/Glebegor/do-app"
	handler "github.com/Glebegor/do-app/pkg/handler"
	repository "github.com/Glebegor/do-app/pkg/repository"
	service "github.com/Glebegor/do-app/pkg/service"
	godotenv "github.com/joho/godotenv"
	_ "github.com/lib/pq"

	logrus "github.com/sirupsen/logrus"
	viper "github.com/spf13/viper"
)

//  @title Todo APP API
//  @version 1.0
// 	@description Server(API, back-end) for Todo Application

// @host localhost:8000
// @BasePath /

//		@securityDefinitions.apikey ApiKeyAuth
//		@in header
//	 @name Authorization

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	if err := initConfig(); err != nil {
		logrus.Fatalf("Error initializing config: %s", err.Error())
	}

	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variables: %s", err.Error())
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
		logrus.Fatalf("failed to initializate db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(todo.Server)
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			logrus.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
		}
	}()
	logrus.Printf("%s: Server is running on port 8000. BTW", time.Now())

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Printf("%s: Server Shutting Down. Press F... Bro... \n", time.Now())

	if err := server.Shutdown(context.Background()); err != nil {
		logrus.Errorf("error accured no server shutting down: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("error accured no db connection close: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
