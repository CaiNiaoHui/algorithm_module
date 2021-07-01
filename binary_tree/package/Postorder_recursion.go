package _package

import (
	"fmt"
)

func Postorder_recursion(root *TreeNode) {
	if root == nil {
		return
	}
	Postorder_recursion(root.Left)
	Postorder_recursion(root.Right)
	fmt.Printf("%d ", root.Value)
}