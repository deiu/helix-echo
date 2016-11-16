package helix

import (
	"github.com/labstack/echo"
	"strings"
)

func CORSHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		origin := c.Request().Header.Get("Origin")
		if len(origin) > 0 {
			c.Response().Header().Set("Access-Control-Allow-Origin", origin)
		}
		if len(origin) < 1 {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		}

		corsReqH := c.Request().Header.Get("Access-Control-Request-Headers") // CORS preflight only
		if len(corsReqH) > 0 {
			c.Response().Header().Set("Access-Control-Allow-Headers", corsReqH)
		}
		corsReqM := c.Request().Header.Get("Access-Control-Request-Method") // CORS preflight only
		if len(corsReqM) > 0 {
			c.Response().Header().Set("Access-Control-Allow-Methods", corsReqM)
		} else {
			c.Response().Header().Set("Access-Control-Allow-Methods", strings.Join(methodsAll, ", "))
		}

		// add HSTS
		c.Response().Header().Set("Strict-Transport-Security", "max-age=63072000; includeSubDomains")

		return next(c)
	}
}
