package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepthExample1(t *testing.T) {
	node15 := TreeNode{Val: 15}
	node7 := TreeNode{Val: 7}
	node20 := TreeNode{Val: 20, Left: &node15, Right: &node7}
	node9 := TreeNode{Val: 9}
	root := TreeNode{Val: 3, Left: &node9, Right: &node20}

	actual := maxDepth(&root)
	assert.Equal(t, 3, actual)
}
func TestMaxDepthExample2(t *testing.T) {
	node2 := TreeNode{Val: 2}
	root := TreeNode{Val: 1, Right: &node2}

	actual := maxDepth(&root)
	assert.Equal(t, 2, actual)
}
func TestMaxDepthNoRoot(t *testing.T) {
	actual := maxDepth(nil)
	assert.Equal(t, 0, actual)
}
func TestMaxDepthOneNode(t *testing.T) {
	root := TreeNode{Val: 3}
	actual := maxDepth(&root)
	assert.Equal(t, 1, actual)
}
