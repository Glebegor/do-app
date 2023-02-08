package main

import (
	"fmt"
	"log"

	todo "github.com/Glebegor/do-app"
)

func main() {
	server := new(todo.Server)
	fmt.Printf("SERVER IS RUNNING. ")
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error while running server: %s", err.Error())
	}
}
