package http

import (
	"html/template"
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
)

var homePage *template.Template

func createPageTemplate(filenames ...string) *template.Template {
	for i := 0; i < len(filenames); i++ {
		filenames[i] = tmplDirectory + filenames[i]
	}

	tmpl := template.Must(template.ParseFiles(filenames...))
	tmpl.ParseGlob(tmplDirectory + "components/*.html")
	return tmpl
}

type pages struct {
	todoService *todo.TodoService
}

func (s *Server) registerPages() {
	homePage = createPageTemplate("layouts/base.html", "pages/home.html")

	p := pages{
		todoService: s.todoService,
	}

	s.echo.GET("/", p.homePage)
}

func (p *pages) homePage(ctx echo.Context) error {
	todos, err := p.todoService.GetTodos()
	if err != nil {
		return nil
	}

	return render(ctx, http.StatusOK, homePage, map[string]any{
		"PageTitle": "Home page",
		"Todos":     todos,
	})
}
