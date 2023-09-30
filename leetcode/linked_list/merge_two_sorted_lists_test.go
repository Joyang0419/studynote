package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/*
Time complexity: O(n + m)
Space complexity:  O(1)

解題思路:
如果 list1 是空的，則返回 list2，因為沒有其他元素需要合併。
如果 list2 是空的，則返回 list1，由於同樣的理由。
比較和遞歸（Comparison and Recursion）:

比較 list1 和 list2 的當前節點的值。
如果 list1 的值小於 list2 的值，
那麼 list1 的當前節點將是合併列表的當前節點。
然後遞歸地調用 mergeTwoLists 函數，
將 list1 的下一個節點和 list2 作為參數，並將結果設置為 list1 的下一個節點。
*/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Val < list2.Val {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}

	list2.Next = mergeTwoLists(list2.Next, list1)
	return list2
}

//func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
//	result := &ListNode{}
//	current := result
//	for list1 != nil && list2 != nil {
//		if list1.Val > list2.Val {
//			current.Next = list2
//			list2 = list2.Next
//		} else {
//			current.Next = list1
//			list1 = list1.Next
//		}
//		if current.Next != nil {
//			current = current.Next
//		}
//	}
//
//	if list1 == nil {
//		current.Next = list2
//	} else {
//		current.Next = list1
//	}
//
//	return result.Next
//}

func TestMergeTwoLists(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	list2 := &ListNode{
		Val: 3,
		Next: &ListNode{
			Val:  4,
			Next: nil,
		},
	}
	result := mergeTwoLists(list1, list2)

	var ints []int
	for result != nil {
		ints = append(ints, result.Val)
		result = result.Next
	}

	assert.ElementsMatch(t, ints, []int{1, 2, 3, 4})

}
