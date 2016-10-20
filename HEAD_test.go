package helix

import (
	// "os"
	// "strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHEADHandler(t *testing.T) {
	assert.NoError(t, HeadHandler(nil))
}

func TestHEADRequest(t *testing.T) {

}
