package helix

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"

	"github.com/labstack/echo/engine/standard"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func init() {
	// uncomment for extra logging
	// os.Setenv("HELIX_LOGGING", "true")
	e := NewServer()
	std := standard.WithTLS("127.0.0.1", "cert.pem", "key.pem")
	std.SetHandler(e)

	// testServer
	testServer = httptest.NewTLSServer(std.Handler)
	testServer.URL = strings.Replace(testServer.URL, "127.0.0.1", "localhost", 1)
	// testClient
	testClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}
}
