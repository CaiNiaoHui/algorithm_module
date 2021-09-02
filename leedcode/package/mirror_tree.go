package _package

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */


func MirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root2 := &TreeNode{
		Val: root.Val,
		Left: root.Right,
		Right: root.Left,
	}

	root2.Left = MirrorTree(root.Right)
	root2.Right = MirrorTree(root.Left)

	return root2
}


