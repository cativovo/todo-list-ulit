package http

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, statusCode int, t templ.Component) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ctx.Response().Writer.WriteHeader(statusCode)
	return t.Render(ctx.Request().Context(), ctx.Response().Writer)
}
