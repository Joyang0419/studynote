package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
Time complexity: O(n)
Space complexity:  O(n)

解題思路:
找尋中間的node
count: 紀錄node的count
用一本map紀錄count, 跟 *node
最後用count / 2 就可以找到中間的數字，再從map去拿出node, 結案
*/
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[int]*ListNode)
	count := 1

	for current := head; current != nil; current = current.Next {
		m[count] = current
		count++
	}

	quotient := count / 2
	if count%2 != 0 {
		quotient = quotient + 1
	}

	return m[quotient]
}

func TestMiddleNode(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  2,
			Next: nil,
		},
	}
	result := middleNode(node)

	assert.Equal(t, result.Val, 2)
}
