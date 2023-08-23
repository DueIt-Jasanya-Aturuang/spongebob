package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRandomChar(t *testing.T) {
	randomChar, err := RandomChar(6)
	assert.NoError(t, err)
	t.Log(randomChar)
	assert.Equal(t, 6, len(randomChar))
}
