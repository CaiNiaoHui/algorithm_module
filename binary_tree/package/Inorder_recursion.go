package _package

import (
	"fmt"
)

func Inorder_recursion(root *TreeNode) {
	if root == nil {
		return
	}
	Inorder_recursion(root.Left)
	fmt.Printf("%d ", root.Value)
	Inorder_recursion(root.Right)
}