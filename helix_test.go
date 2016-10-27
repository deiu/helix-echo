package helix

import (
	"crypto/tls"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/labstack/echo/test"
	"github.com/stretchr/testify/assert"
)

func Test_GET_HTTP1_1(t *testing.T) {
	// Create a temporary http/1.1 client
	httpClient := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
				NextProtos:         []string{"http/1.1"},
			},
		},
	}
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	res, err := httpClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, res.ProtoAtLeast(1, 1))
}

func Test_GET_HTTP2(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	res, err := testClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, res.ProtoAtLeast(2, 0))
}

func Test_Server_RedirectHTTPS(t *testing.T) {
	e := echo.New()
	next := func(c echo.Context) (err error) {
		return c.NoContent(http.StatusOK)
	}
	req := test.NewRequest(echo.GET, "http://example.org", nil)
	res := test.NewResponseRecorder()
	c := e.NewContext(req, res)
	middleware.HTTPSRedirect()(next)(c)
	assert.Equal(t, http.StatusMovedPermanently, res.Status())
	assert.Equal(t, "https://example.org", res.Header().Get(echo.HeaderLocation))
}
