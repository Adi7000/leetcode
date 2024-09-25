package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, 3)
	expected := []int{5, 6, 7, 1, 2, 3, 4}
	assert.Equal(t, expected, nums)
}
func TestRotateArrayLengthEqualRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, 7)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, nums)
}
func TestRotateNoRotate(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	Rotate(nums, 0)
	expected := []int{1, 2, 3, 4, 5, 6, 7}
	assert.Equal(t, expected, nums)
}
func TestRotateArrayLengthOne(t *testing.T) {
	nums := []int{1}
	Rotate(nums, 7)
	expected := []int{1}
	assert.Equal(t, expected, nums)
}
