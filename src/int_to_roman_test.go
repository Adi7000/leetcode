package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRomanExample2(t *testing.T) {
	num := 58
	actual := IntToRoman(num)
	assert.Equal(t, "LVIII", actual)

}
