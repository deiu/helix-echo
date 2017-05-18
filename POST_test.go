package helix

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_POST_HandlerNoRDF(t *testing.T) {
	req, err := http.NewRequest("POST", testServer.URL, nil)
	assert.NoError(t, err)
	resp, err := testClient.Do(req)
	assert.NoError(t, err)
	println("Status code:", resp.StatusCode)
}

func Test_POST_HandlerRDF(t *testing.T) {
	req, err := http.NewRequest("POST", testServer.URL, nil)
	assert.NoError(t, err)
	req.Header.Add("Content-Type", "text/turtle")
	resp, err := testClient.Do(req)
	assert.NoError(t, err)
	println("Status code:", resp.StatusCode)
}

func Test_POST_Request(t *testing.T) {

}
