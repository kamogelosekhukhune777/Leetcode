package easy

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var prevValue = math.MinInt64

var minDiff = math.MaxInt64

func inorderTraversal(node *TreeNode) {
	if node == nil {
		return
	}

	inorderTraversal(node.Left)

	if prevValue != math.MinInt64 {
		minDiff = min(minDiff, node.Val-prevValue)
	}

	prevValue = node.Val

	inorderTraversal(node.Right)
}

func GetMinimumDifference(root *TreeNode) int {
	inorderTraversal(root)
	return minDiff
}
