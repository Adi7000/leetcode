package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoveDuplicates(t *testing.T) {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := removeDuplicates(nums)
	expected := []int{1, 1, 2, 2, 3}

	assert.Equal(t, k, len(expected))
	for i := 0; i < k; i++ {
		assert.Equal(t, expected[i], nums[i])
	}
}

func TestRemoveDuplicatesSmallArray(t *testing.T) {
	nums := []int{1}
	k := removeDuplicates(nums)
	expected := []int{1}

	assert.Equal(t, k, len(expected))
	for i := 0; i < k; i++ {
		assert.Equal(t, expected[i], nums[i])
	}
}
