package main

import (
	"fmt"
	_package "github.com/CaiNiaoHui/algorithm_module/link_list/package"
)

func main() {
	list1 := &_package.ListNode{
		Next:  &_package.ListNode{
			Next:  &_package.ListNode{
				Next:  &_package.ListNode{
					Next:  &_package.ListNode{
						Next:  &_package.ListNode{
							Next:  nil,
							Value: 4,
						},
						Value: 3,
					},
					Value: 3,
				},
				Value: 3,
			},
			Value: 2,
		},
		Value: 1,
	}
	list2 := &_package.ListNode{
		Next:  &_package.ListNode{
			Next:  &_package.ListNode{
				Next:  &_package.ListNode{
					Next:  &_package.ListNode{
						Next:  &_package.ListNode{
							Next:  nil,
							Value: 4,
						},
						Value: 3,
					},
					Value: 3,
				},
				Value: 2,
			},
			Value: 2,
		},
		Value: 1,
	}
	list4 := &_package.ListNode{
		Next:  &_package.ListNode{
			Next:  &_package.ListNode{
				Next:  &_package.ListNode{
					Next:  &_package.ListNode{
						Next:  nil,
						Value: 6,
					},
					Value: 2,
				},
				Value: 12,
			},
			Value: 4,
		},
		Value: 11,
	}

	// 遍历原始链表
	fmt.Println("遍历原始链表list1 list2: ")
	_package.Print_list(list1)
	_package.Print_list(list2)

	// 合并两个升序链表
	fmt.Println("合并两个链表: ")
	list3 := _package.MergeTwoList(list1, list2)
	_package.Print_list(list3)

	// 遍历链表
	fmt.Println("删除重复链表值打印：list1")
	_package.Print_list(list1)
	list1 = _package.DeleteDuplicates(list1)
	_package.Print_list(list1)

	// 反转链表
	fmt.Println("反转链表打印：list2")
	_package.Print_list(list2)
	list2 = _package.ReverseList(list2)
	_package.Print_list(list2)

	// 删除链表中所有重复的值
	fmt.Println("删除所有重复链表值打印：list2")
	_package.Print_list(list2)
	list2 = _package.DeleteDuplicatesAll(list2)
	_package.Print_list(list2)

	// 反转m到n的链表
	fmt.Println("反转m到n的链表打印：list1")
	_package.Print_list(list1)
	list1 = _package.ReverseMN(list1, 2, 3)
	_package.Print_list(list1)

	// 提取l2中大于2的数，拼接
	fmt.Println("拼接l2中大于2的数打印：list4")
	_package.Print_list(list4)
	list4 = _package.Partition(list4, 7)
	_package.Print_list(list4)

	// 归并排序 列表
	fmt.Println("归并排序列表：list2")
	_package.Print_list(list2)
	list2 = _package.SortList(list2)
	_package.Print_list(list2)

	// 判断链表是否有环
	fmt.Println("判断是否有环：list3")
	fmt.Println(_package.HasCycle(list3))

	list5 := &_package.ListNode{
		Next:  &_package.ListNode{
			Next:  &_package.ListNode{
				Next:  &_package.ListNode{
					Next:  &_package.ListNode{
						Next:  nil,
						Value: 1,
					},
					Value: 2,
				},
				Value: 3,
			},
			Value: 2,
		},
		Value: 1,
	}

	// 判断链表是否是回文链表
	fmt.Println("判断是否是回文链表链表：list5")
	fmt.Println(_package.IsPalidrome(list5))

	// 深拷贝
	fmt.Println("深拷贝链表l6：list6")
	list6 := &_package.ListNodeR{
		Next:   &_package.ListNodeR{
			Next:   &_package.ListNodeR{
				Next:   &_package.ListNodeR{
					Next:   &_package.ListNodeR{
						Next:   &_package.ListNodeR{
							Next:   nil,
							Value:  3,
							Random: nil,
						},
						Value:  2,
						Random: nil,
					},
					Value:  1,
					Random: nil,
				},
				Value:  2,
				Random: nil,
			},
			Value:  3,
			Random: nil,
		},
		Value:  4,
		Random: nil,
	}
	list7 := _package.CopyRandomList(list6)
	_package.Print_listR(list7)

	// 深拷贝
	fmt.Println("深拷贝链表l6：list5")
	_package.Print_list(list5)
	list8 := _package.Deep_link_my(list5)
	_package.Print_list(list8)

}

