package helix

import (
	"github.com/labstack/echo"
)

func CORSHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		origin := c.Request().Header().Get("Origin")
		if len(origin) > 0 {
			c.Response().Header().Set("Access-Control-Allow-Origin", origin)
		}
		if len(origin) < 1 {
			c.Response().Header().Set("Access-Control-Allow-Origin", "*")
		}
		return nil
	}
}
