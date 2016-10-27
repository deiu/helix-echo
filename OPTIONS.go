package helix

import (
	"net/http"

	"github.com/labstack/echo"
)

// OptionsHandler uses a closure with the signature func(http.ResponseWriter,
// *http.Request). It sets extra headers that are needed for the CORS preflight
// requests.
func OptionsHandler(c echo.Context) error {

	// Do not return content
	return c.NoContent(http.StatusOK)
}
