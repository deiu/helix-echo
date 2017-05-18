package helix

import (
	"crypto/tls"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Logger_NewLoggerConfig(t *testing.T) {
	conf := NewHelixConfig()
	conf.Logfile = "helix.log"

	logFile := conf.GetLogger()

	o := reflect.TypeOf(logFile).String()
	assert.Equal(t, "*os.File", o)
	stat, err := os.Stat(conf.Logfile)
	assert.NoError(t, err)
	assert.Equal(t, int64(0), stat.Size())

	e := NewServer(conf)
	server := httptest.NewTLSServer(e)
	server.URL = strings.Replace(server.URL, "127.0.0.1", "localhost", 1)
	// testClient
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	request, err := http.NewRequest("GET", server.URL, nil)
	assert.NoError(t, err)
	response, err := client.Do(request)
	assert.NoError(t, err)
	assert.Equal(t, 200, response.StatusCode)

	stat, err = os.Stat(conf.Logfile)
	assert.NoError(t, err)
	assert.NotEqual(t, 0, stat.Size())

	// cleanup
	os.Remove(conf.Logfile)
}
