package helix

import (
	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

var (
	HELIX_VERSION = "0.1"
	methodsAll    = []string{
		"OPTIONS", "HEAD", "GET",
		"PATCH", "POST", "PUT", "MKCOL", "DELETE",
		"COPY", "MOVE", "LOCK", "UNLOCK",
	}
)

// NewLoggerConfig formats the log to a specific template
// func NewLoggerConfig() middleware.LoggerConfig {
// 	return middleware.LoggerConfig{
// 		Format: `[${time_rfc3339}] ${method} request for: ${uri} completed in (${latency_human})` + "\n" +
// 			`[${time_rfc3339}] From: ${remote_ip}` + "\n" +
// 			`[${time_rfc3339}] Status: ${status}` + "\n" +
// 			`[${time_rfc3339}] Bytes in: ${bytes_in}` + "\n" +
// 			`[${time_rfc3339}] Bytes out: ${bytes_out}` + "\n",
// 		Output: os.Stdout,
// 	}
// }

// NewServer creates a new server handler
func NewServer(conf *HelixConfig) *echo.Echo {
	handler := echo.New()

	// Utility Middleware
	// enable logging (change later)
	if conf.Debug {
		handler.Use(middleware.Logger())
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
	handler.GET("/_stats", s.Handle)
	handler.GET("/_info", ServerInfo)
	handler.File("/empty.txt", "empty.txt")

	// CRUD Middleware
	handler.OPTIONS("/*", OptionsHandler)
	handler.HEAD("/*", HeadHandler)
	handler.GET("/*", GetHandler)
	handler.POST("/*", PostHandler)
	handler.PUT("/*", PutHandler)
	handler.DELETE("/*", DeleteHandler)

	return handler
}
