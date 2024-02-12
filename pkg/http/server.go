package http

import (
	"html/template"
	"io"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Server struct {
	echo        *echo.Echo
	todoService *todo.TodoService
}

type TemplateRenderer struct {
	templates *template.Template
}

func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

var renderer = &TemplateRenderer{
	templates: template.Must(template.ParseGlob("frontend/templates/**/*.html")),
}

func NewServer(ts *todo.TodoService) *Server {
	e := echo.New()
	e.Renderer = renderer
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
