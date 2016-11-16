package helix

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"

	"golang.org/x/net/http2"
	// "github.com/labstack/echo"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func init() {
	// uncomment for extra logging
	conf := NewHelixConfig()
	e := NewServer(conf)
	// s := &http.Server{
	// 	TLSConfig: &tls.Config{
	// 		MinVersion:               tls.VersionTLS12,
	// 		NextProtos:               []string{"h2"},
	// 		PreferServerCipherSuites: true,
	// 		CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
	// 		CipherSuites: []uint16{
	// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
	// 			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
	// 			tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
	// 			tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
	// 			tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
	// 			tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	// 		},
	// 	},
	// 	Addr: "127.0.0.1",
	// }

	// std := e.WithTLS("127.0.0.1", "cert.pem", "key.pem")
	// std.SetHandler(e)

	// testServer
	testServer = httptest.NewTLSServer(e)
	testServer.TLS.NextProtos = []string{"h2"}
	testServer.TLS.MinVersion = tls.VersionTLS12
	testServer.TLS.PreferServerCipherSuites = true
	testServer.TLS.CurvePreferences = []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256}
	testServer.TLS.CipherSuites = []uint16{
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		tls.TLS_ECDHE_RSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_RSA_WITH_AES_128_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
		tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
	}
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
