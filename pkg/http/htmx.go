package http

import "github.com/labstack/echo/v4"

func htmxBoosted(ctx echo.Context) bool {
	return ctx.Request().Header.Get("HX-Boosted") == "true"
}
