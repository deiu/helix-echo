package helix

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_MW_ServerInfo(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL+"/_info", nil)
	assert.NoError(t, err)

	res, err := testClient.Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Contains(t, string(body), "Protocol:")
}

func Test_MW_ServerStats(t *testing.T) {
	// Setup
	req, err := http.NewRequest("GET", testServer.URL+"/_stats", nil)
	assert.NoError(t, err)
	res, err := testClient.Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	assert.Equal(t, http.StatusOK, res.StatusCode)
	bbody := []byte(string(body))
	var parsed map[string]interface{}
	err = json.Unmarshal(bbody, &parsed)
	assert.NoError(t, err)
	assert.NotEmpty(t, parsed["requestCount"])
	assert.NotEmpty(t, parsed["statuses"])
	assert.NotEmpty(t, parsed["uptime"])
}
