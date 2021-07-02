package main

import (
	"fmt"

	"github.com/CaiNiaoHui/algorithm_module/binary_tree/package"
)

func init() {

}

func main() {
	// 实例结构体
	tree := &_package.TreeNode{
		Value: 3,
		Left:  &_package.TreeNode{
			Value: 1,
			Left:  nil,
			Right: &_package.TreeNode{
				Value: 2,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &_package.TreeNode{
			Value: 6,
			Left:  nil,
			Right: &_package.TreeNode{
				Value: 9,
				Left:  nil,
				Right: nil,
			},
		},
	}
	// 递归先序
	fmt.Printf("输出递归先序遍历: ")
	_package.Preorder_recursion(tree)
	// 非递归先序
	result := _package.Preorder_non_recursive(tree)
	fmt.Printf("输出非递归先序遍历: ")
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	// 递归中序
	fmt.Printf("输出递归中序遍历: ")
	_package.Inorder_recursion(tree)
	// 非递归中序
	result = _package.Inorder_non_recursive(tree)
	fmt.Printf("输出非递归中序遍历: ")
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	// 递归后序
	fmt.Printf("输出递归后序遍历: ")
	_package.Postorder_recursion(tree)
	// 非递归后序
	result = _package.Postorder_non_recursive(tree)
	fmt.Printf("输出非递归后序遍历: ")
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	// dfs遍历树
	fmt.Printf("dfs遍历树: ")
	result = _package.Dfs_upTodown(tree)
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// dfs遍历树(从下到上)
	fmt.Printf("分治法遍历树: ")
	result = _package.Dfs_downToUp(tree)
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// bfs层序遍历
	fmt.Printf("bfs层序遍历树: ")
	result1 := _package.Bfs_levelOrder(tree)
	for _, list := range result1 {
		for _, val := range list {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	// 归并排序
	arr := []int{3,5,1,2,5,3,2,1,3,14,5,25,435,36,46,341,2}
	result = _package.MergeSort(arr)
	fmt.Printf("归并排序: ")
	for _, val := range result {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// 快速排序
	arr1 := arr
	_ = _package.QuickSort(arr1)
	fmt.Printf("快速排序: ")
	for _, val := range arr1 {
		fmt.Printf("%d ", val)
	}
	fmt.Println()

	// 树的最大深度
	depth := _package.MaxDepth(tree)
	fmt.Printf("树的最大深度: %d", depth)
	fmt.Println()

	// 是否是平衡树
	depth = _package.IsBalaced(tree)
	fmt.Printf("是否是平衡树: %d", depth)
	fmt.Println()

	// 遍历最大路径
	maxpath := _package.MaxPathSum(tree)
	fmt.Printf("树中最长一条路径: %d", maxpath)
	fmt.Println()


	// bfs层序遍历
	fmt.Printf("bfs层序从上到下遍历树: ")
	result1 = _package.Bfs_Traverse(tree)
	for _, list := range result1 {
		for _, val := range list {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	// bfs层序遍历从下到上
	fmt.Printf("bfs层序从下到上遍历树: ")
	result1 = _package.Bfs_Traverse_reverse(tree)
	for _, list := range result1 {
		for _, val := range list {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()

	// bfs层序遍历从下到上
	fmt.Printf("bfs层序Z字形遍历树: ")
	result1 = _package.Bfs_Traverse_Z(tree)
	for _, list := range result1 {
		for _, val := range list {
			fmt.Printf("%d ", val)
		}
	}
	fmt.Println()
	// bfs应用二叉搜索树
	fmt.Printf("是否是二叉排序树: ")
	tr := _package.Is_BST(tree)
	fmt.Println(tr)


}
