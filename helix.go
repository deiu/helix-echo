package helix

import (
	"os"

	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
)

var (
	methodsAll = []string{
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
func NewServer() *echo.Echo {
	handler := echo.New()

	// Utility Middleware
	// enable logging (change later)
	if len(os.Getenv("HELIX_LOGGING")) > 0 {
		handler.Use(middleware.Logger())
	}
	// Server header
	handler.Use(ServerHeader)
	// recover from panics
	handler.Use(middleware.Recover())
	// server stats
	// CORS
	handler.Use(CORSHandler)

	// Routes Middleware
	// HTTP/2 test routes
	handler.GET("/test/info", testRequestInfo)

	// Stats
	s := NewStats()
	handler.Use(s.StatsMiddleware)
	handler.GET("/_stats", s.Handle)

	// CRUD Middleware
	handler.OPTIONS("/*", OptionsHandler)
	handler.HEAD("/*", HeadHandler)
	handler.GET("/*", GetHandler)
	handler.POST("/*", PostHandler)
	handler.PUT("/*", PutHandler)
	handler.DELETE("/*", DeleteHandler)

	return handler
}
