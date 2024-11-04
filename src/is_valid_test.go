package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidExample1(t *testing.T) {
	input := "()"
	actual := isValid(input)
	assert.Equal(t, true, actual)
}
func TestIsValidExample2(t *testing.T) {
	input := "()[]{}"
	actual := isValid(input)
	assert.Equal(t, true, actual)
}
func TestIsValidExample3(t *testing.T) {
	input := "(]"
	actual := isValid(input)
	assert.Equal(t, false, actual)
}
func TestIsValidExample4(t *testing.T) {
	input := "([])"
	actual := isValid(input)
	assert.Equal(t, true, actual)
}
func TestIsValidDanglingBracket(t *testing.T) {
	input := "([])["
	actual := isValid(input)
	assert.Equal(t, false, actual)
}
func TestIsValidNoBrackets(t *testing.T) {
	input := ""
	actual := isValid(input)
	assert.Equal(t, true, actual)
}
func TestIsValidOneClosingBrackets(t *testing.T) {
	input := "]"
	actual := isValid(input)
	assert.Equal(t, false, actual)
}
