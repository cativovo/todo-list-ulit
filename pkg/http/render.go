package http

import (
	"github.com/a-h/templ"
	"github.com/labstack/echo/v4"
)

func Render(c echo.Context, statusCode int, t templ.Component) error {
	c.Response().Header().Set(echo.HeaderContentType, echo.MIMETextHTML)
	c.Response().Writer.WriteHeader(statusCode)
	return t.Render(c.Request().Context(), c.Response().Writer)
}
