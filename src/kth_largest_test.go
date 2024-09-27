package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest1(t *testing.T) {
	nums := []int{3, 2, 1, 5, 6, 4}
	k := 2
	actual := findKthLargest(nums, k)
	assert.Equal(t, 5, actual)
}

func TestFindKthLargest2(t *testing.T) {
	nums := []int{3, 2, 3, 1, 2, 4, 5, 5, 6}
	k := 4
	actual := findKthLargest(nums, k)
	assert.Equal(t, 4, actual)
}
func TestFindKthLargestOneElement(t *testing.T) {
	nums := []int{3}
	k := 1
	actual := findKthLargest(nums, k)
	assert.Equal(t, 3, actual)
}
func TestFindKthLargestSameMaxMin(t *testing.T) {
	nums := []int{3, 3}
	k := 2
	actual := findKthLargest(nums, k)
	assert.Equal(t, 3, actual)
}
