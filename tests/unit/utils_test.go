package unit

import (
	"testing"

	"github.com/DueIt-Jasanya-Aturuang/spongebob/internal/utils"
	"github.com/stretchr/testify/assert"
)

func TestRandomChar(t *testing.T) {
	randomChar, err := utils.RandomChar(6)
	assert.NoError(t, err)
	t.Log(randomChar)
	assert.Equal(t, 6, len(randomChar))
}
