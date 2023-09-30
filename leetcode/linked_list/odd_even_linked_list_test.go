package linked_list

/*
tc: O(n)
sc: O(1)

desc:
兩根指針: odd, even
evenHead 保留
创建了两个指针 odd 和 even 分别指向初始的奇数节点和偶数节点。
通过一个循环，不断地更新奇数节点和偶数节点的 Next 指针，
使得奇数节点指向下一个奇数节点，偶数节点指向下一个偶数节点。

在循环结束后，将奇数链表的最后一个节点的 Next 指针指向偶数链表的头节点，从而连接两个链表。
最后返回原链表的头节点 head，它现在已经是按照奇偶顺序重新排列过的链表。
*/
func oddEvenList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	odd := head       // 奇數指針
	even := head.Next // 偶數指針
	evenHead := even  // 保存偶数链表的头节点，用于最后的连接

	// 当偶数节点和偶数节点的下一个节点存在时，进行循环
	for even != nil && even.Next != nil {
		odd.Next = even.Next // 使奇数节点指向下一个奇数节点
		odd = odd.Next       // 更新奇数节点
		even.Next = odd.Next // 使偶数节点指向下一个偶数节点
		even = even.Next     // 更新偶数节点
	}

	odd.Next = evenHead // 连接奇数链表和偶数链表
	return head
}
