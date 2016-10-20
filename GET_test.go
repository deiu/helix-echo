package helix

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	// "os"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/engine/standard"
	// "github.com/linkeddata/helix"
	"github.com/stretchr/testify/assert"
)

var (
	testServer *httptest.Server
	testClient *http.Client
)

func init() {
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

func TestGETHandler(t *testing.T) {
	assert.NoError(t, GetHandler(nil))
}

func TestGETServerInfo(t *testing.T) {
	// Setup
	req, err := http.NewRequest(echo.GET, testServer.URL+"/test/info", nil)
	assert.NoError(t, err)
	res, err := testClient.Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	// Assertions
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.NotEmpty(t, body)
}
