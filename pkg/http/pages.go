package http

import (
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
)

type pages struct {
	todoService *todo.TodoService
}

func (s *Server) registerPages() {
	p := pages{
		todoService: s.todoService,
	}

	s.echo.GET("/", p.homePage)
}

func (p *pages) homePage(c echo.Context) error {
	todos, err := p.todoService.GetTodos()
	if err != nil {
		return nil
	}

	return c.Render(http.StatusOK, "home.html", todos)
}
