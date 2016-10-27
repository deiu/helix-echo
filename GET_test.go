package helix

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GET_Request(t *testing.T) {
	req, err := http.NewRequest("GET", testServer.URL, nil)
	assert.NoError(t, err)
	res, err := testClient.Do(req)
	assert.NoError(t, err)
	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	assert.Equal(t, "Hello world", string(body))
}

func BenchmarkGET(b *testing.B) {
	e := 0
	for i := 0; i < b.N; i++ {
		req, _ := http.NewRequest("GET", testServer.URL, nil)
		if res, _ := testClient.Do(req); res.StatusCode != 200 {
			e++
		}
	}
	if e > 0 {
		b.Log(fmt.Sprintf("%d/%d failed", e, b.N))
		b.Fail()
	}
}
