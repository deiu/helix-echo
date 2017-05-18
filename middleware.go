// Extra middleware
package helix

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"sync"
	"time"

	"github.com/labstack/echo"
	"github.com/rs/zerolog"
)

type (
	Stats struct {
		Uptime       time.Time      `json:"uptime"`
		RequestCount uint64         `json:"requestCount"`
		Statuses     map[string]int `json:"statuses"`
		mutex        sync.RWMutex
	}
)

func NewStats() *Stats {
	return &Stats{
		Uptime:   time.Now(),
		Statuses: make(map[string]int),
	}
}

// Process is the middleware function.
func (s *Stats) StatsMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		if err := next(c); err != nil {
			c.Error(err)
		}
		s.mutex.Lock()
		defer s.mutex.Unlock()
		s.RequestCount++
		status := strconv.Itoa(c.Response().Status)
		s.Statuses[status]++
		return nil
	}
}

func ServerInfo(c echo.Context) error {
	req := c.Request()
	format := "\n<pre>\nProtocol: %s\nHost: %s\nMethod: %s\nPath: %s\n</pre>\n"
	return c.HTML(http.StatusOK, fmt.Sprintf(format, req.Proto, req.Host, req.Method, req.URL))
}

// Handle is the endpoint to get stats.
func (s *Stats) Handler(c echo.Context) error {
	s.mutex.RLock()
	defer s.mutex.RUnlock()
	return c.JSON(http.StatusOK, s)
}

// ServerHeader middleware sets a Server header to the response.
func ServerHeader(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set(echo.HeaderServer, "Helix-"+HELIX_VERSION)
		return next(c)
	}
}

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		logger = zerolog.New(os.Stderr).With().
			Timestamp().
			Str("Method", c.Request().Method).
			Str("Path", c.Request().URL.String()).
			Logger()

		return next(c)
	}
}
