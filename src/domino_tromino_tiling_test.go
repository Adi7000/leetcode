package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumTilingsDominos(t *testing.T) {
	assert.Equal(t, 1, numTilingsDominos(1))
	assert.Equal(t, 2, numTilingsDominos(2))
	assert.Equal(t, 3, numTilingsDominos(3))
	assert.Equal(t, 5, numTilingsDominos(4))
	assert.Equal(t, 8, numTilingsDominos(5))
}
