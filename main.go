package main

import (
	"github.com/cativovo/todo-list-ulit/pkg/http"
	repository "github.com/cativovo/todo-list-ulit/pkg/repository/memory"
	"github.com/cativovo/todo-list-ulit/pkg/todo"
)

func main() {
	memoryRepository := repository.NewMemoryRepository()
	todoService := todo.NewTodoService(memoryRepository)

	server := http.NewServer(todoService)
	server.ListenAndServe("127.0.0.1:4000")
}
