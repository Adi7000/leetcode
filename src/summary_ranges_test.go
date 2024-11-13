package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSummaryRangesExample1(t *testing.T) {
	nums := []int{0, 1, 2, 4, 5, 7}
	actual := summaryRanges(nums)
	expected := []string{"0->2", "4->5", "7"}
	assert.Equal(t, expected, actual)
}
func TestSummaryRangesExample2(t *testing.T) {
	nums := []int{0, 2, 3, 4, 6, 8, 9}
	actual := summaryRanges(nums)
	expected := []string{"0", "2->4", "6", "8->9"}
	assert.Equal(t, expected, actual)
}
func TestSummaryRangesNoInput(t *testing.T) {
	nums := []int{}
	actual := summaryRanges(nums)
	expected := []string{}
	assert.Equal(t, expected, actual)
}
func TestSummaryRangesSingleInput(t *testing.T) {
	nums := []int{0}
	actual := summaryRanges(nums)
	expected := []string{"0"}
	assert.Equal(t, expected, actual)
}
