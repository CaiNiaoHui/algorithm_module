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
```

## how to use algorithm?

Examples:

```go
binary_tree := &tree.TreeNode{
	Value: 1,
	Left:  &tree.TreeNode{
		Value: 2,
		Left:  nil,
		Right: &tree.TreeNode{
			Value: 3,
			Left:  nil,
			Right: nil,
		},
	},
	Right: nil,
}

result := tree.Bfs_levelOrder(binary_tree)

// bfs层序遍历
fmt.Printf("bfs层序遍历树: ")
for _, list := range result {
	for _, val := range list {
		fmt.Printf("%d ", val)
	}
}
fmt.Println()
```
