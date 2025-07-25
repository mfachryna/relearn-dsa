package relearndsa

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	if root.Left == nil && root.Right == nil {
		return 1
	}

	left := 0
	if root.Left != nil {
		left += maxDepth(root.Left) + 1
	}

	right := 0
	if root.Right != nil {
		right += maxDepth(root.Right) + 1
	}

	return max(right, left)
}
