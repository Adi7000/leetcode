package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTrapExample1(t *testing.T) {
	heights := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	actual := trap(heights)
	assert.Equal(t, 6, actual)
}
func TestTrapExample2(t *testing.T) {
	heights := []int{4, 2, 0, 3, 2, 5}
	actual := trap(heights)
	assert.Equal(t, 9, actual)
}
