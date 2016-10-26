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

var (
	methodsAll = []string{
		"OPTIONS", "HEAD", "GET",
		"PATCH", "POST", "PUT", "MKCOL", "DELETE",
		"COPY", "MOVE", "LOCK", "UNLOCK",
	}
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
// 	ticker := timhandler.NewTicker(1 * timhandler.Second)
// 	defer ticker.Stop()

// 	for {
// 		fmt.Fprintf(res, "%v\n", timhandler.Now())
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
	handler := echo.New()

	// Utility Middleware
	// enable logging (change later)
	if len(os.Getenv("HELIX_LOGGING")) > 0 {
		handler.Use(middleware.Logger())
	}
	// recover from panics
	handler.Use(middleware.Recover())
	// CORS
	handler.Use(CORSHandler)

	// Routes Middleware
	// HTTP/2 test routes
	handler.GET("/test/info", requestInfo)
	// handler.GET("/test/stream", streamTime)
	// CRUD Middleware
	handler.OPTIONS("/*", OptionsHandler)
	handler.HEAD("/*", HeadHandler)
	handler.GET("/*", GetHandler)
	handler.POST("/*", PostHandler)
	handler.PUT("/*", PutHandler)
	handler.DELETE("/*", DeleteHandler)

	return handler
}
