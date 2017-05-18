package helix

import (
	"crypto/tls"
	"errors"
	"github.com/labstack/echo"

	"github.com/labstack/echo/middleware"
	"github.com/rs/zerolog"
)

var (
	HELIX_VERSION = "0.1"
	methodsAll    = []string{
		"OPTIONS", "HEAD", "GET", "POST", "PUT", "PATCH", "DELETE",
	}
	logger zerolog.Logger
)

// NewServer creates a new server handler
func NewServer(conf *HelixConfig) *echo.Echo {
	handler := echo.New()

	// ****** Utility Middleware ******
	// enable logging (change later)

	if !conf.Debug {
		zerolog.SetGlobalLevel(zerolog.Disabled)
	}
	// handler.Use(middleware.LoggerWithConfig(NewLoggerConfig(conf.Logfile)))
	logger = zerolog.New(conf.GetLogger()).With().Timestamp().Logger()

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
	handler.Use(LoggerMiddleware)
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

func NewTLSConfig(cert, key string) (*tls.Config, error) {
	var err error
	cfg := &tls.Config{}

	if len(cert) == 0 || len(key) == 0 {
		return cfg, errors.New("Missing cert and key for TLS configuration")
	}

	cfg.MinVersion = tls.VersionTLS12
	cfg.NextProtos = []string{"h2"}
	// use strong crypto
	cfg.PreferServerCipherSuites = true
	cfg.CurvePreferences = []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}
	cfg.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	}
	cfg.Certificates = make([]tls.Certificate, 1)
	cfg.Certificates[0], err = tls.LoadX509KeyPair(cert, key)

	return cfg, err
}
