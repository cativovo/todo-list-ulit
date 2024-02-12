package http

import (
	"html/template"
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
)

var (
	homePage *template.Template
	todoPage *template.Template
)

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
	todoPage = createPageTemplate("layouts/base.html", "pages/todo.html")

	p := pages{
		todoService: s.todoService,
	}

	s.echo.GET("/", p.homePage)
	s.echo.GET("/:id", p.todoPage)
}

func (p *pages) homePage(ctx echo.Context) error {
	todos, err := p.todoService.GetTodos()
	data := map[string]any{
		"PageTitle": "Home page",
		"Todos":     todos,
		"Err":       err != nil,
	}

	if err != nil {
		return render(ctx, http.StatusInternalServerError, homePage, data)
	}

	return render(ctx, http.StatusOK, homePage, data)
}

func (p *pages) todoPage(ctx echo.Context) error {
	id := ctx.Param("id")
	todo, err := p.todoService.GetTodo(id)
	boosted := htmxBoosted(ctx)
	data := map[string]any{
		"PageTitle": "Todo page",
		"Todo":      todo,
		"Boosted":   boosted,
		"Err":       err != nil,
	}

	var tmpl *template.Template
	if boosted {
		tmpl = todoPage.Lookup("content")
	} else {
		tmpl = todoPage
	}

	if err != nil {
		return render(ctx, http.StatusInternalServerError, tmpl, data)
	}

	return render(ctx, http.StatusOK, tmpl, data)
}
