package main

import (
	"log"
)

func main() {
	server := new(todo.Server)
	if err := server.Run("8000"); err != nil {
		log.Fatalf("Error while running server")
	}
}
