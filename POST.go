package helix

import (
	// "github.com/deiu/rdf2go"
	"github.com/labstack/echo"
)

func PostHandler(c echo.Context) error {
	if gotRDF(c.Request().Header.Get("Content-Type")) {
		return PostHandleRDF(c)
	}

	return nil
}

func PostHandleRDF(c echo.Context) error {
	c.Response().Status = 201
	return nil
}
