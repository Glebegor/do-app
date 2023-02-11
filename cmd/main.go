package main

import (
	"fmt"
	"log"
	"time"

	 "github.com/Glebegor/do-app"
	 "github.com/Glebegor/do-app/pkg/handler"
)

func main() {
	handlers := new(handler.Handler) 
	server := new(todo.Server)

	fmt.Printf("%s: Server is running on port 8000. BTW", time.Now())
	if err := server.Run("8000", handlers.initRoutes()); err != nil {
		log.Fatalf("Error while running server: %s, %s", err.Error(), time.Now())
	}
}
