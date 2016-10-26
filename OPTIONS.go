package helix

import ()

import (
	"github.com/labstack/echo"
)

// OptionsHandler uses a closure with the signature func(http.ResponseWriter,
// *http.Request). It sets extra headers that are needed for the CORS preflight
// requests.
func OptionsHandler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return nil
	}
}
