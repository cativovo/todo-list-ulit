package http

import (
	"html/template"

	"github.com/labstack/echo/v4"
)

func render(ctx echo.Context, statusCode int, tmpl *template.Template, data any) error {
	ctx.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	ctx.Response().Writer.WriteHeader(statusCode)
	return tmpl.Execute(ctx.Response().Writer, data)
}
