package _package

import (
	"fmt"
)

func Preorder_recursion(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%d ", root.Value)
	Preorder_recursion(root.Left)
	Preorder_recursion(root.Right)
}