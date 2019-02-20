package user

import (
	"testing"

	"gotest.tools/assert"
)

func TestFail(t *testing.T) {
	assert.Equal(t, 0, 1)
}
