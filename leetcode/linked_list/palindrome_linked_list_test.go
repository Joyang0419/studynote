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
for left, right 從左往中間走，跟右邊往中間走, 去比對，有一個不一樣，就代表不是Palindrome
*/
func isPalindrome(head *ListNode) bool {
	if head == nil {
		return false
	}

	m := make(map[int]*ListNode)
	count := 1

	for current := head; current != nil; current = current.Next {
		m[count] = current
		count++
	}

	for left, right := 1, count-1; left < right; left, right = left+1, right-1 {
		if m[left].Val != m[right].Val {
			return false
		}
	}

	return true
}

func TestIsPalindrome(t *testing.T) {
	node := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val:  1,
			Next: nil,
		},
	}
	result := isPalindrome(node)

	assert.Equal(t, result, true)
}
