package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

/*
time complexity: O(n)，其中 n 是鏈表中的節點數量。

space complexity: 空間複雜度是由算法所需的額外空間決定的。
在這個算法中，我們只使用了幾個額外的變量 (prev, curr, 和 nextTemp)，
並且沒有使用任何額外的數據結構，如堆栈或隊列。
因此，空間複雜度是 O(1)，這意味著它是一個常數空間複雜度，與輸入大小無關。

解題思路:
curr 就是一根指針, 靠著替換curr 變成nextTemp 去迭代更換
*/
func reverseList(head *ListNode) *ListNode {
	var prev *ListNode

	curr := head
	for curr != nil {
		nextTemp := curr.Next // 要先存起來

		curr.Next = prev
		prev = curr

		curr = nextTemp // 迭代原本的每個ListNode
	}

	return prev
}

func TestReverseList(t *testing.T) {
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
	result := reverseList(node)

	var ints []int
	for result != nil {
		ints = append(ints, result.Val)
		result = result.Next
	}

	assert.ElementsMatch(t, ints, []int{3, 2, 1})
}
