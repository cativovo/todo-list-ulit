package http

import (
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/cativovo/todo-list-ulit/web/layouts"
	pageTempl "github.com/cativovo/todo-list-ulit/web/pages"
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
	s.echo.GET("/todo/:id", p.todoPage)
}

func (p *pages) homePage(ctx echo.Context) error {
	pageTitle := "Home"
	todos, err := p.todoService.GetTodos()
	if err != nil {
		return render(ctx, http.StatusInternalServerError, layouts.Base(pageTempl.Error(), pageTitle, false))
	}

	homeProps := pageTempl.HomeProps{
		Todos: todos,
		Title: pageTitle,
	}

	return render(ctx, http.StatusOK, layouts.Base(pageTempl.Home(homeProps), pageTitle, htmxBoosted(ctx)))
}

func (p *pages) todoPage(ctx echo.Context) error {
	pageTitle := "Todo"
	id := ctx.Param("id")
	todo, err := p.todoService.GetTodo(id)
	if err != nil {
		ctx.Logger().Error(err)
		return render(ctx, http.StatusInternalServerError, layouts.Base(pageTempl.Error(), pageTitle, false))
	}

	todoProps := pageTempl.TodoProps{
		Todo: todo,
	}

	return render(ctx, http.StatusOK, layouts.Base(pageTempl.Todo(todoProps), pageTitle, htmxBoosted(ctx)))
}
