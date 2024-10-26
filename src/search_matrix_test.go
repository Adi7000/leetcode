package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchMatrixExample1(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	actual := searchMatrix(matrix, 3)
	assert.Equal(t, true, actual)
}
func TestSearchMatrixExample2(t *testing.T) {
	matrix := [][]int{
		{1, 3, 5, 7},
		{10, 11, 16, 20},
		{23, 30, 34, 60},
	}
	actual := searchMatrix(matrix, 13)
	assert.Equal(t, false, actual)
}
func TestSearchMatrixEmpty(t *testing.T) {
	matrix := [][]int{}
	actual := searchMatrix(matrix, 9)
	assert.Equal(t, false, actual)
}
func TestSearchMatrixOneRow(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}}
	actual := searchMatrix(matrix, 5)
	assert.Equal(t, true, actual)
}
