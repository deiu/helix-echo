package helix

import (
	"io/ioutil"
	"net/http"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestOPTIONSWithoutReqMethod(t *testing.T) {
// 	request, err := http.NewRequest("OPTIONS", testServer.URL, nil)
// 	assert.NoError(t, err)
// 	// request.Header.Add("Access-Control-Request-Method", "PATCH")
// 	response, err := testClient.Do(request)
// 	assert.NoError(t, err)
// 	body, err := ioutil.ReadAll(response.Body)
// 	response.Body.Close()
// 	assert.Empty(t, string(body))
// 	assert.Equal(t, 200, response.StatusCode)
// 	assert.Equal(t, strings.Join(methodsAll, ", "), response.Header.Get("Access-Control-Allow-Methods"))
// }

func TestOPTIONSWithReqMethod(t *testing.T) {
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
