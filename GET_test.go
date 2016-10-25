package helix

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestGETEmptyHandler(t *testing.T) {
	assert.NoError(t, GetHandler(nil))
}

func TestGETCORS(t *testing.T) {
	// Setup
	req, err := http.NewRequest(echo.GET, testServer.URL, nil)
	assert.NoError(t, err)
	t.Run("Empty Origin", func(t *testing.T) {
		res, err := testClient.Do(req)
		assert.NoError(t, err)

		// Assertions
		assert.Equal(t, "*", res.Header.Get("Access-Control-Allow-Origin"))
	})
	t.Run("Set Origin", func(t *testing.T) {
		exUri := "https://example.org"
		req.Header.Set("Origin", exUri)
		res, err := testClient.Do(req)
		assert.NoError(t, err)

		// Assertions
		assert.Equal(t, exUri, res.Header.Get("Access-Control-Allow-Origin"))
	})
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
