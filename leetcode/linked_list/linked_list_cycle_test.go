package linked_list

/*
解題思路: 使用快慢指針的方法（也稱為Floyd的循環檢測算法）來解決。
*/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	slow := head
	fast := head.Next

	for fast != nil && fast.Next != nil {
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}

	return false
}
