package http

import (
	"html/template"
	"net/http"

	"github.com/cativovo/todo-list-ulit/pkg/todo"
	"github.com/labstack/echo/v4"
)

type handlers struct {
	todoService *todo.TodoService
}

var componentsTmpl = template.Must(template.ParseGlob(tmplDirectory + "components/*.html"))

func (s *Server) registerHandlers() {
	h := handlers{
		todoService: s.todoService,
	}

	s.echo.POST("/todo", h.addTodo)
	s.echo.PATCH("/todo/:id", h.updateTodo)
	s.echo.DELETE("/todo/:id", h.deleteTodo)
}

func (h *handlers) addTodo(ctx echo.Context) error {
	t := todo.Todo{
		Title:       ctx.FormValue("title"),
		Description: ctx.FormValue("description"),
	}

	todo, err := h.todoService.AddTodo(t)
	if err != nil {
		// TODO: handle
		return nil
	}

	ctx.Response().Header().Set(HXRetarget, "#todos")
	ctx.Response().Header().Set(HXReswap, "afterbegin")
	return render(ctx, http.StatusOK, componentsTmpl.Lookup("todo"), todo)
}

func (h *handlers) updateTodo(ctx echo.Context) error {
	t := todo.Todo{
		ID:          ctx.Param("id"),
		Title:       ctx.FormValue("title"),
		Description: ctx.FormValue("description"),
	}

	todo, err := h.todoService.UpdateTodo(t)
	if err != nil {
		ctx.Logger().Error(err)
		// TODO: handle
		return nil
	}

	data := todoPageProps{
		PageTitle: "Todo page",
		Todo:      todo,
		Boosted:   false,
		Err:       err != nil,
	}

	ctx.Response().Header().Set(HXTrigger, "refetchTodo")
	return render(ctx, http.StatusNoContent, nil, data)
}

func (h *handlers) deleteTodo(ctx echo.Context) error {
	id := ctx.Param("id")

	err := h.todoService.DeleteTodo(id)
	if err != nil {
		// TODO: handle
		return nil
	}

	if ctx.Request().Header.Get("Redirect") == "true" {
		htmxLocation(ctx, "/")
	}

	// https://htmx.org/attributes/hx-delete/
	return ctx.NoContent(http.StatusOK)
}
