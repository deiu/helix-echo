package helix

import (
	"io/ioutil"
	"net/http"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	exUri = "https://example.org"
)

func Test_CORS_NoOrigin(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)
	res, err := testClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, "*", res.Header.Get("Access-Control-Allow-Origin"))
}

func Test_CORS_WithOrigin(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)

	req.Header.Set("Origin", exUri)
	res, err := testClient.Do(req)
	assert.NoError(t, err)

	assert.Equal(t, exUri, res.Header.Get("Access-Control-Allow-Origin"))
}
func Test_CORS_AllowHeaders(t *testing.T) {
	request, err := http.NewRequest("OPTIONS", testServer.URL, nil)
	assert.NoError(t, err)
	request.Header.Add("Access-Control-Request-Headers", "User, ETag")
	response, err := testClient.Do(request)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.Empty(t, string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "User, ETag", response.Header.Get("Access-Control-Allow-Headers"))
}

func Test_CORS_WithoutReqMethod(t *testing.T) {
	request, err := http.NewRequest("OPTIONS", testServer.URL, nil)
	assert.NoError(t, err)
	// request.Header.Add("Access-Control-Request-Method", "PATCH")
	response, err := testClient.Do(request)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.Empty(t, string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, strings.Join(methodsAll, ", "), response.Header.Get("Access-Control-Allow-Methods"))
}

func Test_CORS_WithReqMethod(t *testing.T) {
	request, err := http.NewRequest("OPTIONS", testServer.URL, nil)
	assert.NoError(t, err)
	request.Header.Add("Access-Control-Request-Method", "PATCH")
	response, err := testClient.Do(request)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(response.Body)
	response.Body.Close()
	assert.Empty(t, string(body))
	assert.Equal(t, 200, response.StatusCode)
	assert.Equal(t, "PATCH", response.Header.Get("Access-Control-Allow-Methods"))
}
