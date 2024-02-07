package http

import (
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	partialsTemplates "github.com/cativovo/todo-list-ulit/templates/partials"
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
	s.echo.PUT("/todo/:id", h.updateTodo)
	s.echo.DELETE("/todo/:id", h.deleteTodo)
}

func (h *handlers) addTodo(c echo.Context) error {
	t := todo.Todo{
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	todo, err := h.todoService.AddTodo(t)
	if err != nil {
		return Render(c, http.StatusInternalServerError, partialsTemplates.Error())
	}

	return Render(c, http.StatusOK, partialsTemplates.TodoItem(todo))
}

func (h *handlers) updateTodo(c echo.Context) error {
	t := todo.Todo{
		ID:          c.Param("id"),
		Title:       c.FormValue("title"),
		Description: c.FormValue("description"),
	}

	todo, err := h.todoService.UpdateTodo(t)
	if err != nil {
		return Render(c, http.StatusInternalServerError, partialsTemplates.Error())
	}

	return Render(c, http.StatusOK, partialsTemplates.TodoItem(todo))
}

func (h *handlers) deleteTodo(c echo.Context) error {
	id := c.Param("id")

	err := h.todoService.DeleteTodo(id)
	if err != nil {
		return Render(c, http.StatusInternalServerError, partialsTemplates.Error())
	}

	// https://htmx.org/attributes/hx-delete/
	return c.NoContent(http.StatusOK)
}
