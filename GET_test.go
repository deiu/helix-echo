package helix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GET_EmptyHandler(t *testing.T) {
	assert.NoError(t, GetHandler(nil))
}
