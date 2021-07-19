### link_list function use

1. create list
   ```cgo
   list := &_package.ListNode{Value: 0, Next: nil}
   ```

2. list traversal
    ```gotemplate
    _package.Print_list(list1)
    ```
   
3. merge link-list
    ```gotemplate
    list3 := _package.MergeTwoList(list1, list2)
    ```
   
4. Delete duplicate 
    ```gotemplate
    list1 = _package.DeleteDuplicates(list1)
    ```
   
5.  reverse list 
    ```gotemplate
    list2 = _package.ReverseList(list2)
    ```

6. Delete all duplicates 
    ```gotemplate
    list2 = _package.DeleteDuplicatesAll(list2)
    ```
  
7. Extract the number greater than 2 in l2, and splice
    ```gotemplate
    list4 = _package.Partition(list4, 7)
    ```

8. Reverse list from M to N
    ```gotemplate
    list1 = _package.ReverseMN(list1, 2, 3)
    ```
   
9. 归并排序
    ```gotemplate
    list2 = _package.SortList(list2)
    ```
   
10. 判断是否有环
    ```gotemplate
    _package.HasCycle(list3
    ```
    
11. 判断回文链表
    ```gotemplate
    _package.IsPalidrome(list5)
    ```
    
12. 深拷贝
    ```gotemplate
    list6 = _package.CopyRandomList(list6)
    ```
    
