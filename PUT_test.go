package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PUT_Handler(t *testing.T) {
	assert.NoError(t, PutHandler(nil))
}

func Test_PUT_Request(t *testing.T) {

}
