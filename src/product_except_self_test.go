package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductExceptSelfExample1(t *testing.T) {
	nums := []int{1, 2, 3, 4}
	actual := productExceptSelf(nums)
	expected := []int{24, 12, 8, 6}
	assert.Equal(t, expected, actual)
}
func TestProductExceptSelfExample2(t *testing.T) {
	nums := []int{-1, 1, 0, -3, 3}
	actual := productExceptSelf(nums)
	expected := []int{0, 0, 9, 0, 0}
	assert.Equal(t, expected, actual)
}
func TestProductExceptSelfNoNums(t *testing.T) {
	nums := []int{}
	actual := productExceptSelf(nums)
	assert.Nil(t, actual)
}
func TestProductExceptSelfOneNum(t *testing.T) {
	nums := []int{1}
	actual := productExceptSelf(nums)
	assert.Nil(t, actual)
}
