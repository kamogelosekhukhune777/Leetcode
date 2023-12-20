package easy

//defined in 530-minimum-absolute-difference-in-a-bst
/*
type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
*/

//SearchBST instead of searchBST
func SearhBST(root *TreeNode, val int) *TreeNode {
	if root == nil {
		return nil
	}

	if root.Val == val {
		return root
	} else if root.Val < val {
		return SearhBST(root.Right, val)
	} else {
		return SearhBST(root.Left, val)
	}
}
