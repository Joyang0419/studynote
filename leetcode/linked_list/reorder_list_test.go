package linked_list

import (
	"fmt"
	"testing"
)

/*
tc: O(n)
sc: O(n)

desc:
*/
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil {
		return
	}

	/*
		創建一個空的節點數組：
		首先，我們創建了一個空的節點數組 nodes，然後遍歷了整個鏈表，將每個節點的引用添加到該數組中。
	*/
	var nodes []*ListNode
	for node := head; node != nil; node = node.Next {
		nodes = append(nodes, node)
	}

	/*
		使用雙指針技術：
		接下來，我們使用兩個指針 i 和 j，它們分別從數組的開頭和結尾開始。 i 指針從前向後移動，而 j 指針從後向前移動。
	*/
	i, j := 0, len(nodes)-1
	for i < j {
		/*
			重新連接節點：
			我們持續進行以下操作，直到 i 和 j 相遇或者 i 超過 j：
			將 nodes[i] 的下一個節點設置為 nodes[j]。
			增加 i 的值。
			如果 i 和 j 相遇，則跳出循環。
			將 nodes[j] 的下一個節點設置為 nodes[i]。
			減少 j 的值。
		*/
		nodes[i].Next = nodes[j]
		i++
		if i == j {
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}

	// 結束連接：
	// 最後，我們將 nodes[i] 的 Next 指針設置為 nil，以確保列表在末尾結束，而不是形成一個循環。
	nodes[i].Next = nil
}

func TestReorderList(t *testing.T) {
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
	reorderList(node)

	var ints []int
	for node != nil {
		ints = append(ints, node.Val)
		node = node.Next
	}

	for _, num := range ints {
		fmt.Println(num)
	}

}
