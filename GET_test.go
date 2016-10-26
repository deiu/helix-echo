package helix

import (
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_GET_HTTP2(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	res, err := testClient.Do(req)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)
	assert.True(t, res.ProtoAtLeast(2, 0))
}

func Test_GET_EmptyHandler(t *testing.T) {
	assert.NoError(t, GetHandler(nil))
}

func Test_GET_ServerInfo(t *testing.T) {
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
