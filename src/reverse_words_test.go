package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWordsExample1(t *testing.T) {
	s := "the sky is blue"
	expected := "blue is sky the"

	assert.Equal(t, ReverseWords(s), expected)
}

func TestReverseWordsExample2(t *testing.T) {
	s := "  hello   world  "
	expected := "world hello"

	assert.Equal(t, ReverseWords(s), expected)
}
