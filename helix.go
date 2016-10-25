package helix

import (
	"fmt"
	"net/http"
	"os"
	// "time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	"github.com/labstack/echo/middleware"
)

//----------
// Test Handlers
//----------

func requestInfo(c echo.Context) error {
	req := c.Request().(*standard.Request).Request
	format := "\nProtocol: %s\nHost: %s\nRemote Address: %s\nMethod: %s\nPath: %s\n\n"
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.RemoteAddr, req.Method, req.URL))
}

// func streamTime(c echo.Context) error {
// 	res := c.Response().(*standard.Response).ResponseWriter
// 	gone := res.(http.CloseNotifier).CloseNotify()
// 	res.Header().Set(echo.HeaderContentType, "text/turtle")
// 	res.WriteHeader(http.StatusOK)
// 	ticker := time.NewTicker(1 * time.Second)
// 	defer ticker.Stop()

// 	for {
// 		fmt.Fprintf(res, "%v\n", time.Now())
// 		res.(http.Flusher).Flush()
// 		select {
// 		case <-ticker.C:
// 		case <-gone:
// 			break
// 		}
// 	}
// }

// NewServer creates a new server handler
func NewServer() *echo.Echo {
	e := echo.New()

	// Utility Middleware
	// enable logging (change later)
	if len(os.Getenv("HELIX_LOGGING")) > 0 {
		e.Use(middleware.Logger())
	}
	// recover from panics
	e.Use(middleware.Recover())
	// CORS
	e.Use(CORSHandler)

	// Routes Middleware
	// HTTP/2 test routes
	e.GET("/test/info", requestInfo)
	// e.GET("/test/stream", streamTime)
	// CRUD Middleware
	e.OPTIONS("/*", HeadHandler)
	e.HEAD("/*", HeadHandler)
	e.GET("/*", GetHandler)
	e.POST("/*", PostHandler)
	e.PUT("/*", PutHandler)
	e.DELETE("/*", DeleteHandler)

	return e
}
