package http

import (
	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo        *echo.Echo
	todoService *todo.TodoService
}

func NewServer(ts *todo.TodoService) *Server {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	s := &Server{
		echo:        e,
		todoService: ts,
	}

	s.registerHandlers()
	s.registerPages()

	return s
}

func (s *Server) ListenAndServe(addr string) {
	if err := s.echo.Start(addr); err != nil {
		s.echo.Logger.Fatal(err)
	}
}
