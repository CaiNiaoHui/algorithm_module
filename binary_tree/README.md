## 知识点

### 二叉树遍历

**前序遍历**：**先访问根节点**，再前序遍历左子树，再前序遍历右子树
**中序遍历**：先中序遍历左子树，**再访问根节点**，再中序遍历右子树
**后序遍历**：先后序遍历左子树，再后序遍历右子树，**再访问根节点**

注意点

- 以根访问顺序决定是什么遍历
- 左子树都是优先右子树

#### 前序递归

- watch "Preorder_recursion.go"

#### 前序非递归

- watch "Preorder_non_recursive.go"

#### 中序非递归

- watch "In_order_non_recursive.go"

#### 后序非递归

- watch "Postorder_non_recursive"

#### DFS 深度搜索-从上到下

- 实质就是树的先序遍历算法
- watch "dfs_upTodown"


#### 分治法模版

```
type ResultType []int

func traversal(root *TreeNode) ResultType {
	//nil
	result := make(ResultType, 0)
	if root == nil {
		return  nil
	}
	// Divide
	ResultType left := traversal(root.Left)
	ResultType right := traversal(root.Right)

	//Conquer
	ResultType result = Merge from left and right

	return result
}

```
