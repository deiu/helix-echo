package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDELETEHandler(t *testing.T) {
	assert.NoError(t, DeleteHandler(nil))
}

func TestDELETERequest(t *testing.T) {

}
