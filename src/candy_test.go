package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCandyExample1(t *testing.T) {
	nums := []int{1, 0, 2}
	actual := candy(nums)
	assert.Equal(t, 5, actual)
}
func TestCandyExample2(t *testing.T) {
	nums := []int{1, 2, 2}
	actual := candy(nums)
	assert.Equal(t, 4, actual)
}
func TestCandyLocalMaximaHasOneGap(t *testing.T) {
	nums := []int{1, 3, 2}
	actual := candy(nums)
	assert.Equal(t, 4, actual)
}
func TestCandyLocalMaximaHasManyGaps(t *testing.T) {
	nums := []int{1, 2, 3, 2, 1}
	actual := candy(nums)
	assert.Equal(t, 9, actual)
}
func TestCandyNoChildren(t *testing.T) {
	nums := []int{}
	actual := candy(nums)
	assert.Equal(t, 0, actual)
}
