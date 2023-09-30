package linked_list

/*
tc: O(n log n)
sc: O(log n)

desc:
sortList 函数是主要的函数，它负责将链表分成两半，分别对它们进行排序，然后将它们合并在一起。
findMiddle 函数使用快慢指针方法来找到链表的中点。
merge 函数将两个已排序的链表合并成一个新的已排序的链表。
*/
func sortList(head *ListNode) *ListNode {
	// Base case: if the list has 0 or 1 node
	if head == nil || head.Next == nil {
		return head
	}

	// Find the middle of the list
	mid := findMiddle(head)
	right := mid.Next
	mid.Next = nil // Split the list into two halves

	// Recursively sort each half
	left := sortList(head)
	right = sortList(right)

	// Merge the sorted halves
	return merge(left, right)
}

func findMiddle(head *ListNode) *ListNode {
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func merge(left, right *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for left != nil && right != nil {
		if left.Val < right.Val {
			current.Next = left
			left = left.Next
		} else {
			current.Next = right
			right = right.Next
		}
		current = current.Next
	}
	if left != nil {
		current.Next = left
	} else {
		current.Next = right
	}
	return dummy.Next
}
