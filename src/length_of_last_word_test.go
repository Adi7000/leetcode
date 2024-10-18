package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLengthOfLastWordExample1(t *testing.T) {
	sentence := "Hello World"
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 5, actual)
}
func TestLengthOfLastWordExample2(t *testing.T) {
	sentence := "   fly me   to   the moon  "
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 4, actual)
}
func TestLengthOfLastWordExample3(t *testing.T) {
	sentence := "luffy is still joyboy"
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 6, actual)
}
func TestLengthOfLastWordNoWords(t *testing.T) {
	sentence := ""
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 0, actual)
}
func TestLengthOfLastOnlySpaces(t *testing.T) {
	sentence := "  "
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 0, actual)
}
func TestLengthOfLastOneWordOnly(t *testing.T) {
	sentence := "Tomy"
	actual := lengthOfLastWord(sentence)
	assert.Equal(t, 4, actual)
}
