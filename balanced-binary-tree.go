package relearndsa

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	diff := abs(countDepth(root.Left) - countDepth(root.Right))

	if diff <= 1 {
		return isBalanced(root.Left) && isBalanced(root.Right)
	}

	return false
}

func countDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	left := countDepth(root.Left)
	right := countDepth(root.Right)

	return max(left, right) + 1
}

func abs(val int) int {
	if val >= 0 {
		return val
	}

	return val * -1
}
