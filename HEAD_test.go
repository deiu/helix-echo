package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HEAD_Handler(t *testing.T) {
	assert.NoError(t, HeadHandler(nil))
}

func Test_HEAD_Request(t *testing.T) {

}
