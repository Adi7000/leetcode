package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHIndexExample1(t *testing.T) {
	nums := []int{3, 0, 6, 1, 5}
	actual := hIndex(nums)
	assert.Equal(t, 3, actual)
}

func TestHIndexExample2(t *testing.T) {
	nums := []int{1, 2, 1}
	actual := hIndex(nums)
	assert.Equal(t, 1, actual)
}

func TestHIndexNoPapers(t *testing.T) {
	nums := []int{}
	actual := hIndex(nums)
	assert.Equal(t, 0, actual)
}
func TestHIndexNoCitationsEqualHIndex(t *testing.T) {
	nums := []int{0, 4, 0, 4}
	actual := hIndex(nums)
	assert.Equal(t, 2, actual)
}
func TestHIndexOnePaperLessThanTotal(t *testing.T) {
	nums := []int{0}
	actual := hIndex(nums)
	assert.Equal(t, 0, actual)
}
func TestHIndexOnePaperMoreThanTotal(t *testing.T) {
	nums := []int{100}
	actual := hIndex(nums)
	assert.Equal(t, 1, actual)
}
