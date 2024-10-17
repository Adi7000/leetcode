package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelfExample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	actual := productExceptSelfMemOptimized(nums)
	expected := []int{24, 12, 8, 6}
	assert.Equal(t, expected, actual)
}
func TestProductExceptSelfExample2(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	actual := productExceptSelfMemOptimized(nums)
	expected := []int{0, 0, 9, 0, 0}
	assert.Equal(t, expected, actual)
}
func TestProductExceptSelfNonOneBoundaryValues(t *testing.T) {
	nums := []int{4, 2, 3, 4}
	actual := productExceptSelfMemOptimized(nums)
	expected := []int{24, 48, 32, 24}
	assert.Equal(t, expected, actual)
}
func TestProductExceptSelfNoNums(t *testing.T) {
	nums := []int{}
	actual := productExceptSelfMemOptimized(nums)
	assert.Nil(t, actual)
}
func TestProductExceptSelfOneNum(t *testing.T) {
	nums := []int{1}
	actual := productExceptSelfMemOptimized(nums)
	assert.Nil(t, actual)
}
