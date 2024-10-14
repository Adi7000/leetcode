package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRobExample1(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	actual := rob(nums)
	assert.Equal(t, 4, actual)
}
func TestRobExample2(t *testing.T) {
	nums := []int{2, 7, 9, 3, 1}
	actual := rob(nums)
	assert.Equal(t, 12, actual)
}
func TestRobNoHouses(t *testing.T) {
	nums := []int{}
	actual := rob(nums)
	assert.Equal(t, 0, actual)
}
func TestRobOneHouse(t *testing.T) {
	nums := []int{1}
	actual := rob(nums)
	assert.Equal(t, 1, actual)
}
func TestRobIncreasingSequence(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5}
	actual := rob(nums)
	assert.Equal(t, 9, actual)
}
