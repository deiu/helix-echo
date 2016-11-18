package helix

import (
	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

var (
	HELIX_VERSION = "0.1"
	methodsAll    = []string{
		"OPTIONS", "HEAD", "GET", "POST", "PUT", "PATCH", "DELETE",
	}
)

// NewServer creates a new server handler
func NewServer(conf *HelixConfig) *echo.Echo {
	handler := echo.New()

	// Utility Middleware
	// enable logging (change later)
	if len(conf.Logfile) > 0 {
		handler.Use(middleware.LoggerWithConfig(NewLoggerWithFile(conf.Logfile)))
	}
	// Server header
	handler.Use(ServerHeader)
	// recover from panics
	handler.Use(middleware.Recover())

	// CORS
	handler.Use(CORSHandler)

	// ****** Routes Middleware ********

	// Server info
	s := NewStats()
	handler.Use(s.StatsMiddleware)
	handler.GET("/_stats", s.Handler)
	handler.GET("/_info", ServerInfo)

	// CRUD Middleware
	handler.OPTIONS("/*", OptionsHandler)
	handler.HEAD("/*", HeadHandler)
	handler.GET("/*", GetHandler)
	handler.POST("/*", PostHandler)
	handler.PUT("/*", PutHandler)
	handler.PATCH("/*", PatchHandler)
	handler.DELETE("/*", DeleteHandler)

	return handler
}
