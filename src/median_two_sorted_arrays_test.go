package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMedianSortedArraysExample1(t *testing.T) {
	nums1 := []int{1, 3}
	nums2 := []int{2}

	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, 2.0, actual)
}
func TestFindMedianSortedArraysExample2(t *testing.T) {
	nums1 := []int{1, 2}
	nums2 := []int{3, 4}

	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, 2.5, actual)
}
func TestFindMedianSortedArrays2MiddleNumsInSameSlice(t *testing.T) {
	nums1 := []int{1, 3, 5}
	nums2 := []int{0}

	actual := findMedianSortedArrays(nums1, nums2)
	assert.Equal(t, 2.0, actual)
}
