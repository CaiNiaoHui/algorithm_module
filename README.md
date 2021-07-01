title: this is the algorithm patter
---

**use algorithm with me!**

step 1 patter :including the Binary tree, Linked list，Stack, Queue and Binary
step 2 leetcode part Q&A

---

- step 1
    1. Binary tree

    2. Linked list

    3. Stack and queue
    
    4. Binary

- step 2
    1. leetcode question
    
---

## Install

```console
go get github.com/CaiNiaoHui/algorithm_module
import (
	tree "github.com/CaiNiaoHui/algorithm_module/binary_tree/package"

) 
```

## how to use algorithm?

Examples:

```go

binary_tree := &tree.TreeNode{
		Value: 1,
		Left:  &tree.TreeNode{
			Value: 2,
			Left:  nil,
			Right: nil,
		},
		Right: nil,
	}

tree.Bfs_levelOrder(Tree)
tree.MaxPathSum(Tree)
```
