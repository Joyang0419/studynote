package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
tc: O(n)
sc: O(1)

desc:
首先检查链表是否为空，或者 k 是否为零，因为在这些情况下不需要旋转。
然后，遍历链表以找到其长度和尾节点。
使用模运算来确定实际的旋转步骤，然后找到新的尾节点。
最后，执行旋转操作：设置新的头节点，断开旧的尾节点，并将旧的尾节点连接到旧的头节点。
*/
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil || k == 0 {
		return head
	}

	// Step 1: Find the length of the list and the tail node
	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	// Step 2: Get the actual rotation steps, and find the new tail
	k %= length
	if k == 0 {
		return head
	}

	newTail := head
	for i := 0; i < length-k-1; i++ {
		newTail = newTail.Next
	}

	// Step 3: Perform the rotation
	newHead := newTail.Next
	newTail.Next = nil // 這時他指的是head

	tail.Next = head

	return newHead
}

func TestRotateRight(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	result := rotateRight(node, 1)

	var ints []int
	for result != nil {
		ints = append(ints, result.Val)
		result = result.Next
	}

	assert.ElementsMatch(t, ints, []int{3, 1, 2})
}
