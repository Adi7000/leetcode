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

func TestRemoveElementNoElement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	actual := removeElement(nums, 7)
	expected := 8
	expectedNums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedNums, nums)
}

func TestRemoveElementFirstELement(t *testing.T) {
	nums := []int{0, 1, 2, 2, 3, 0, 4, 2}
	actual := removeElement(nums, 0)
	expected := 6
	expectedNums := []int{2, 1, 2, 2, 3, 4, 0, 0}
	assert.Equal(t, expected, actual)
	assert.Equal(t, expectedNums, nums)
}
