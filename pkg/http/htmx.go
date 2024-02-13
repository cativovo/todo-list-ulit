package http

import (
	"github.com/labstack/echo/v4"
)

const (
	HXRequest  = "HX-Request"
	HXBoosted  = "HX-Boosted"
	HXReswap   = "HX-Reswap"
	HXRetarget = "HX-Retarget"
	HXTrigger  = "HX-Trigger"
)

func htmxBoosted(ctx echo.Context) bool {
	return ctx.Request().Header.Get(HXBoosted) == "true" || ctx.Request().Header.Get(HXRequest) == "true"
}

func htmxLocation(ctx echo.Context, url string) {
	ctx.Response().Header().Add("HX-Location", url)
}
