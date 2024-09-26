package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanJumpFalse(t *testing.T) {
	nums := []int{3, 2, 1, 0, 4}
	actual := canJump(nums)
	assert.Equal(t, false, actual)
}

func TestCanJumpTrue(t *testing.T) {
	nums := []int{2, 3, 1, 1, 4}
	actual := canJump(nums)
	assert.Equal(t, true, actual)
}

func TestCanJumpOneNum(t *testing.T) {
	nums := []int{0}
	actual := canJump(nums)
	assert.Equal(t, true, actual)
}

func TestCanJumpZeroStart(t *testing.T) {
	nums := []int{0, 1, 2}
	actual := canJump(nums)
	assert.Equal(t, false, actual)
}
