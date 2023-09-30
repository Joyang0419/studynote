package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
tc: O(L)，其中 L 是链表的长度。
sc: 此算法只使用了常量的额外空间，除了输入和输出之外，没有使用任何额外的数据结构。因此，空间复杂度是 O(1)。

desc:
一開始用dummy 是準備就是回傳next, 所以next直接放head, 主要是方便操作整個linked list
fast, slow 就是快慢指針, 這題的關鍵就是讓快慢指針差距是n
然後讓fast走到最後一個位置, 這時slow的位置就是要改next的位置
-> 刪除NthNodeFromEnd -> slow.Next = slow.Next.Next
*/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast := dummy
	slow := dummy

	for i := 0; i <= n; i++ {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	// Removes the nth node from the end
	slow.Next = slow.Next.Next

	return dummy.Next
}

func TestRemoveNthFromEnd(t *testing.T) {
	list1 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val:  5,
						Next: nil,
					},
				},
			},
		},
	}

	result := removeNthFromEnd(list1, 1)
	var ints []int
	for result != nil {
		ints = append(ints, result.Val)
		result = result.Next
	}
	assert.ElementsMatch(t, ints, []int{1, 2, 3, 4})
}
