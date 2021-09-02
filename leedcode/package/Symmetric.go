package _package

/*
对称二叉树

请实现一个函数，用来判断一棵二叉树是不是对称的。如果一棵二叉树和它的镜像一样，那么它是对称的。

例如，二叉树 [1,2,2,3,4,4,3] 是对称的。

    1
   / \
  2   2
 / \ / \
3  4 4  3
但是下面这个 [1,2,2,null,3,null,3] 则不是镜像对称的:

    1
   / \
  2   2
   \   \
   3    3

 

示例 1：

输入：root = [1,2,2,3,4,4,3]
输出：true
示例 2：

输入：root = [1,2,2,null,3,null,3]
输出：false

*/


func IsSymmetric(root *TreeNode) bool  {
	if root == nil {
		return true
	}

	result := recur(root.Left, root.Right)

	return result
}

func recur(l *TreeNode, r *TreeNode) bool {

	if l == nil && r == nil {
		return true
	}
	if (l != nil && r != nil) &&  l.Val == r.Val  {
		result := recur(l.Left, r.Right) && recur(l.Right, r.Left)
		return result
	}
	return false
}
