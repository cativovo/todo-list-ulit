package repository

import (
	"time"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
)

var idCounter = 0

type MemoryRepository struct {
	todos []todo.Todo
}

func NewMemoryRepository() *MemoryRepository {
	mr := &MemoryRepository{}

	// add sample todos
	mr.AddTodo(todo.Todo{Title: "todo title 1", Description: "todo description1"})
	mr.AddTodo(todo.Todo{Title: "todo title 2", Description: "todo description2"})
	mr.AddTodo(todo.Todo{Title: "todo title 3", Description: "todo description3"})
	mr.AddTodo(todo.Todo{Title: "todo title 4", Description: "todo description4"})

	return mr
}

func (mr *MemoryRepository) AddTodo(t todo.Todo) (todo.Todo, error) {
	t.ID = idCounter
	idCounter++
	t.CreatedAt = time.Now()

	mr.todos = append(mr.todos, t)

	return t, nil
}

func (mr *MemoryRepository) GetTodos() ([]todo.Todo, error) {
	return mr.todos, nil
}
