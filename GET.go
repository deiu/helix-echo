package helix

import (
	"net/http"

	"github.com/labstack/echo"
)

func GetHandler(c echo.Context) error {
	return c.HTML(http.StatusOK, "Hello world")
}
