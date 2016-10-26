package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DELETE_Handler(t *testing.T) {
	assert.NoError(t, DeleteHandler(nil))
}

func Test_DELETE_Request(t *testing.T) {

}
