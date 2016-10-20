package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPOSTHandler(t *testing.T) {
	assert.NoError(t, PostHandler(nil))
}

func TestPOSTRequest(t *testing.T) {

}
