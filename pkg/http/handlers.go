package http

import (
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/cativovo/todo-list-ulit/templates"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	todoService *todo.TodoService
}

func (s *Server) registerHandlers() {
	h := handlers{
		todoService: s.todoService,
	}

	s.echo.POST("/todo", h.addTodo)
}

func (h *handlers) addTodo(c echo.Context) error {
	t := todo.Todo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}
	todo, err := h.todoService.AddTodo(t)
	if err != nil {
		return Render(c, http.StatusInternalServerError, templates.Error())
	}

	return Render(c, http.StatusOK, templates.TodoItem(todo))
}
