### tree function use

1. create tree
   ```cgo
   tree := &_package.TreeNode{Value: 0, Left: nil, Right: nil}
   ```

2. Preorder traversal
    ```gotemplate
    _package.Preorder_recursion(tree)
    ```
   
3. Non-recursive preorder traversal
    ```gotemplate
    result := _package.Preorder_non_recursive(tree)
    ```
   
4. Midorder traversal
    ```gotemplate
    _package.Inorder_recursion(tree)
    ```
   
5. Non-recursive Midorder traversal
    ```gotemplate
    result = _package.Inorder_non_recursive(tree)
    ```

6. Postorder traversal
    ```gotemplate
    _package.Postorder_recursion(tree)
    ```
  
7. Non-recursive Postorder traversal
    ```gotemplate
    result = _package.Postorder_non_recursive(tree)
    ```

8. dfs traversal tree (up to down)
    ```gotemplate
    result = _package.Dfs_upTodown(tree)
    ```
   
9. dfs divide and conquer traversing the tree (down to up)
    ```gotemplate
    result = _package.Dfs_downToUp(tree)
    ```
   
10. bfs sequence traversal
    ```gotemplate
    result1 := _package.Bfs_levelOrder(tree)
    ```
    
11. Merge sort
    ```gotemplate
    result = _package.MergeSort(arr)
    ```
    
12. Quick sort
    ```gotemplate
    _ = _package.QuickSort(arr1)
    ```
    
13. tree Maximum depth
    ```gotemplate
    depth := _package.MaxDepth(tree)
    ```
    
14. detect balanced binary tree
    ```gotemplate
    depth = _package.IsBalaced(tree)
    ```
    
15. travel max path
    ```gotemplate
    maxpath := _package.MaxPathSum(tree)
    ```
    
16. bfs sequence traversal (up to down)
    ```gotemplate
    result1 = _package.Bfs_Traverse(tree)
    ```
    
17. bfs sequence traversal (down to up)
    ```gotemplate
    result1 = _package.Bfs_Traverse_reverse(tree)
    ```
    
18. bfs sequence traversal ('Z' glyph traversal)
    ```gotemplate
    result1 = _package.Bfs_Traverse_Z(tree)
    ```
    
19. Judging the binary search tree
    ```gotemplate
    tr := _package.Is_BST(tree)
    ```
