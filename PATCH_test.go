package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_PATCH_Handler(t *testing.T) {
	assert.NoError(t, PatchHandler(nil))
}

func Test_PATCH_Request(t *testing.T) {

}
