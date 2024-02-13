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

type todoModalComponentProps struct {
	ModalTitle     string
	SubmitLabel    string
	ShowModalLabel string
	Method         string
	Url            string
}

type homePageProps struct {
	todoModalComponentProps
	PageTitle string
	Todos     []todo.Todo
	Err       bool
	Boosted   bool
}

func (p *pages) homePage(ctx echo.Context) error {
	todos, err := p.todoService.GetTodos()
	boosted := htmxBoosted(ctx)

	data := homePageProps{
		todoModalComponentProps: todoModalComponentProps{
			ModalTitle:     "Add Todo",
			SubmitLabel:    "Add",
			Method:         "post",
			Url:            "/todo",
			ShowModalLabel: "Add Todo",
		},
		PageTitle: "Home page",
		Todos:     todos,
		Err:       err != nil,
		Boosted:   boosted,
	}

	var tmpl *template.Template
	if boosted {
		tmpl = homePage.Lookup("content")
	} else {
		tmpl = homePage
	}

	if err != nil {
		return render(ctx, http.StatusInternalServerError, tmpl, data)
	}

	return render(ctx, http.StatusOK, tmpl, data)
}

type todoPageProps struct {
	todoModalComponentProps
	PageTitle string
	Todo      todo.Todo
	Boosted   bool
	Err       bool
}

func (p *pages) todoPage(ctx echo.Context) error {
	id := ctx.Param("id")
	todo, err := p.todoService.GetTodo(id)
	boosted := htmxBoosted(ctx)
	data := todoPageProps{
		todoModalComponentProps: todoModalComponentProps{
			ModalTitle:     "Update Todo",
			SubmitLabel:    "Update",
			Method:         "patch",
			Url:            "/todo/" + todo.ID,
			ShowModalLabel: "Update",
		},
		PageTitle: "Todo page",
		Todo:      todo,
		Boosted:   boosted,
		Err:       err != nil,
	}

	var tmpl *template.Template
	if boosted {
		tmpl = todoPage.Lookup("content")
	} else {
		tmpl = todoPage
	}

	if err != nil {
		ctx.Logger().Error(err)
		return render(ctx, http.StatusInternalServerError, tmpl, data)
	}

	return render(ctx, http.StatusOK, tmpl, data)
}
