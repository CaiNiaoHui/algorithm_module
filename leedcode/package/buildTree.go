package _package

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

func BuildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val:   preorder[0],
		Left:  nil,
		Right: nil,
	}

	i := 0
	for ; i < len(preorder); i++ {
		if preorder[0] == inorder[i] {
			break
		}
	}

	root.Left = BuildTree(preorder[1: len(inorder[:i])+1], inorder[:i])
	root.Right = BuildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])

	return root
}
