package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
tc:
sc:

desc:
使用recursive
先設定終止條件，就是傳入的head or head next == nil 終止
然後把firstNode = head, secondNode = head.Next
// 不要想得太複雜, 就是一步一步的去想，就會想通
firstNode.Next = swapPairs(secondNode.Next)
secondNode.Next = firstNode
*/
func swapPairs(head *ListNode) *ListNode {
	// 暫停recursive的條件
	if head == nil || head.Next == nil {
		return head
	}

	// 要交換的node
	firstNode := head
	secondNode := head.Next

	//// Swap the pair of nodes and recursively call the function for the rest of the list
	firstNode.Next = swapPairs(secondNode.Next)
	secondNode.Next = firstNode

	return secondNode
}

func TestSwapPairs(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val:  4,
					Next: nil,
				},
			},
		},
	}
	result := swapPairs(node)

	var ints []int
	for result != nil {
		ints = append(ints, result.Val)
		result = result.Next
	}

	assert.ElementsMatch(t, ints, []int{2, 1, 4, 3})
}
