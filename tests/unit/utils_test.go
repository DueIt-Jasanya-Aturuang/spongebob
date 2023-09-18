package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/helper"
)

func TestRandomChar(t *testing.T) {
	randomChar, err := helper.RandomChar(6)
	assert.NoError(t, err)
	t.Log(randomChar)
	assert.Equal(t, 6, len(randomChar))
}
