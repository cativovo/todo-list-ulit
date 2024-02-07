package repository

import (
	"errors"
	"strconv"
	"time"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
)

var idCounter = 0

type MemoryRepository struct {
	todos map[string]todo.Todo
}

func NewMemoryRepository() *MemoryRepository {
	mr := &MemoryRepository{
		todos: make(map[string]todo.Todo),
	}

	// add sample todos
	mr.AddTodo(todo.Todo{Title: "todo title 1", Description: "todo description1"})
	mr.AddTodo(todo.Todo{Title: "todo title 2", Description: "todo description2"})
	mr.AddTodo(todo.Todo{Title: "todo title 3", Description: "todo description3"})
	mr.AddTodo(todo.Todo{Title: "todo title 4", Description: "todo description4"})

	return mr
}

func (mr *MemoryRepository) AddTodo(t todo.Todo) (todo.Todo, error) {
	t.ID = strconv.Itoa(idCounter)
	idCounter++
	t.CreatedAt = time.Now()
	t.UpdatedAt = time.Now()

	mr.todos[t.ID] = t

	return t, nil
}

func (mr *MemoryRepository) GetTodo(id string) (todo.Todo, error) {
	result, ok := mr.todos[id]
	if !ok {
		return todo.Todo{}, errors.New("todo not found")
	}

	return result, nil
}

func (mr *MemoryRepository) GetTodos() ([]todo.Todo, error) {
	todos := make([]todo.Todo, 0)

	for _, v := range mr.todos {
		todos = append(todos, v)
	}

	return todos, nil
}

func (mr *MemoryRepository) UpdateTodo(t todo.Todo) (todo.Todo, error) {
	newT, err := mr.GetTodo(t.ID)
	if err != nil {
		return todo.Todo{}, err
	}

	newT.Title = t.Title
	newT.Description = t.Description
	newT.UpdatedAt = time.Now()
	mr.todos[t.ID] = newT

	return newT, nil
}

func (mr *MemoryRepository) DeleteTodo(id string) error {
	if _, err := mr.GetTodo(id); err != nil {
		return err
	}

	delete(mr.todos, id)

	return nil
}
