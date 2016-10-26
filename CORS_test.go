package helix

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	exUri = "https://example.org"
)

func TestCORSNoOrigin(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)
	res, err := testClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, "*", res.Header.Get("Access-Control-Allow-Origin"))
}

func TestCORSWithOrigin(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	req.Header.Set("Origin", exUri)
	res, err := testClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, exUri, res.Header.Get("Access-Control-Allow-Origin"))
}
