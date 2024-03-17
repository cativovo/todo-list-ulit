package main

import (
	"fmt"
	"log"
	"os"

	"github.com/cativovo/todo-list-ulit/pkg/http"
	repository "github.com/cativovo/todo-list-ulit/pkg/repository/postgres"
	"github.com/cativovo/todo-list-ulit/pkg/todo"
)

func main() {
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL_MODE"),
	)
	pgRepository, err := repository.NewPostgresRepository(connStr)
	if err != nil {
		log.Fatal(err)
	}

	todoService := todo.NewTodoService(pgRepository)
	server := http.NewServer(todoService)
	server.ListenAndServe("127.0.0.1:4000")
}
