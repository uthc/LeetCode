/* 删除链表的倒数第N个节点
给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？ */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	top, del, fin := head, head, head
	for i := 0; i < n; i++ {
		fin = fin.Next
	}
	if fin == nil {
		return top.Next
	}
	for ; fin.Next != nil; fin = fin.Next {
		del = del.Next
	}
	fin = del.Next
	del.Next = fin.Next
	return top
}