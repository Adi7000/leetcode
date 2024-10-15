package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJumpExample1(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	actual := jump(nums)
	assert.Equal(t, 2, actual)
}
func TestJumpExample2(t *testing.T) {
	nums := []int{2, 3, 0, 1, 4}
	actual := jump(nums)
	assert.Equal(t, 2, actual)
}
func TestJumpNoJump(t *testing.T) {
	nums := []int{2}
	actual := jump(nums)
	assert.Equal(t, 0, actual)
}
func TestEffectiveJump(t *testing.T) {
	nums := []int{2, 3, 4}
	actual := effectiveJump(nums)
	assert.Equal(t, 4, actual)
}
func TestEffectiveJumpFirstElementLarge(t *testing.T) {
	nums := []int{7, 3, 4}
	actual := effectiveJump(nums)
	assert.Equal(t, 5, actual)
}
func TestEffectiveJumpOneElement(t *testing.T) {
	nums := []int{7}
	actual := effectiveJump(nums)
	assert.Equal(t, 7, actual)
}
func TestEffectiveJumpNoElements(t *testing.T) {
	nums := []int{}
	actual := effectiveJump(nums)
	assert.Equal(t, 0, actual)
}
