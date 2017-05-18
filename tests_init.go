package helix

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"strings"

	"golang.org/x/net/http2"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func init() {
	// uncomment for extra logging
	conf := NewHelixConfig()
	conf.Debug = true
	conf.Cert = "test_cert.pem"
	conf.Key = "test_key.pem"

	// prepare TLS config
	tlsCfg, err := NewTLSConfig(conf.Cert, conf.Key)
	if err != nil {
		println(err.Error())
		return
	}

	// testServer
	e := NewServer(conf)
	testServer = httptest.NewUnstartedServer(e)
	testServer.TLS = tlsCfg
	testServer.StartTLS()

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
