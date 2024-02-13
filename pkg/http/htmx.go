package http

import (
	"github.com/labstack/echo/v4"
)

func htmxBoosted(ctx echo.Context) bool {
	return ctx.Request().Header.Get("HX-Boosted") == "true" || ctx.Request().Header.Get("HX-Request") == "true"
}

func htmxLocation(ctx echo.Context, url string) {
	ctx.Response().Header().Add("HX-Location", url)
}
