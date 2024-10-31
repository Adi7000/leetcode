package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCombinationsExample1(t *testing.T) {
	nums := "23"
	actual := letterCombinations(nums)
	expected := []string{
		"ad", "ae", "af",
		"bd", "be", "bf",
		"cd", "ce", "cf",
	}
	assert.Equal(t, expected, actual)
}
func TestLetterCombinationsExample2(t *testing.T) {
	nums := ""
	actual := letterCombinations(nums)
	expected := []string{}
	assert.Equal(t, expected, actual)
}
func TestLetterCombinationsExample3(t *testing.T) {
	nums := "2"
	actual := letterCombinations(nums)
	expected := []string{"a", "b", "c"}
	assert.Equal(t, expected, actual)
}
func TestLetterCombinations4LetterDigit(t *testing.T) {
	nums := "9"
	actual := letterCombinations(nums)
	expected := []string{"w", "x", "y", "z"}
	assert.Equal(t, expected, actual)
}
