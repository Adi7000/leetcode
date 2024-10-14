package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	actual := removeElement(nums, 2)
	expected := 5
	expectedNums := []int{0, 1, 4, 0, 3, 2, 2, 2}
	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedNums, nums)
}
