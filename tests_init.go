package helix

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"

	"golang.org/x/net/http2"

	"github.com/labstack/echo/engine/standard"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func init() {
	// uncomment for extra logging
	conf := NewHelixConfig()
	e := NewServer(conf)
	std := standard.WithTLS("127.0.0.1", "cert.pem", "key.pem")
	std.SetHandler(e)

	// testServer
	testServer = httptest.NewTLSServer(std.Handler)
	testServer.TLS.NextProtos = []string{"h2"}
	testServer.URL = strings.Replace(testServer.URL, "127.0.0.1", "localhost", 1)
	// testClient
	testClient = &http.Client{
		Transport: &http2.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				NextProtos:         []string{"h2"},
			},
		},
	}
}
