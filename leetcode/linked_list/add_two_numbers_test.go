package linked_list

/*
tc:  O(max(n, m))
sc: 空间复杂度为 O(max(n, m)) 或者 O(max(n, m) + 1)
（在最坏的情况下，两个链表的和可能会产生一个新的节点）

desc:
初始化：设置一个 dummy（哑）节点，以及一个当前节点
current 指向 dummy。设置一个变量 carry 来跟踪进位，初始值为 0。

遍历两个链表：同时遍历两个链表，直到两个链表都已遍历完为止。
在每次迭代中，对两个节点的值以及 carry 进行求和。

添加新节点：从总和中获取当前值，
并创建一个新的节点，将其添加到 current 节点的后面。然后更新进位 carry。

处理进位：如果在结束遍历后 carry 仍然不为
（例如两个较长的数字相加可能产生额外的进位），需要添加一个值为 carry 的新节点。

返回结果：返回 dummy 节点的下一个节点，因为 dummy 节点本身不包含有效数据。
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10 // 十進位
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}

	return dummy.Next
}
