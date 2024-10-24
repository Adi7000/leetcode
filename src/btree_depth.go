package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthNonEmpty(root)
}
func maxDepthNonEmpty(root *TreeNode) int {
	if root.Left == nil {
		if root.Right == nil {
			return 1
		} else {
			return 1 + maxDepthNonEmpty(root.Right)
		}
	} else {
		if root.Right == nil {
			return 1 + maxDepthNonEmpty(root.Left)
		} else {
			return 1 + max(maxDepthNonEmpty(root.Right), maxDepthNonEmpty(root.Left))
		}
	}
}
