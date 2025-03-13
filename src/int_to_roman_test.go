package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToRomanExample1(t *testing.T) {
	num := 3749
	actual := IntToRoman(num)
	assert.Equal(t, "MMMDCCXLIX", actual)

}
func TestIntToRomanExample2(t *testing.T) {
	num := 58
	actual := IntToRoman(num)
	assert.Equal(t, "LVIII", actual)

}
func TestIntToRomanExample3(t *testing.T) {
	num := 1994
	actual := IntToRoman(num)
	assert.Equal(t, "MCMXCIV", actual)

}
