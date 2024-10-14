package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMajorityElementExample1(t *testing.T) {
	nums := []int{3, 2, 3}
	actual := majorityElement(nums)
	assert.Equal(t, 3, actual)
}
func TestMajorityElementExample2(t *testing.T) {
	nums := []int{2, 2, 1, 1, 1, 2, 2}
	actual := majorityElement(nums)
	assert.Equal(t, 2, actual)
}
func TestMajorityElementInSecondHalfOnly(t *testing.T) {
	nums := []int{1, 1, 2, 3, 4, 4, 4, 4, 4}
	actual := majorityElement(nums)
	assert.Equal(t, 4, actual)
}
func TestMajorityElementOneElement(t *testing.T) {
	nums := []int{1}
	actual := majorityElement(nums)
	assert.Equal(t, 1, actual)
}
