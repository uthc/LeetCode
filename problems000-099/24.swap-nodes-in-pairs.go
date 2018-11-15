/* 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.
说明:

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	top := head
	if top != nil && top.Next != nil {
		top = top.Next
	}
	p := head
	til := p
	if p != nil {
		a := p
		b := p.Next
		p = p.Next
		if p == nil {
			return top
		}
		//交换a,b
		p = p.Next
		b.Next = a
		a.Next = p
		til = a
	}
	if p != nil {
		til.Next = swapPairs(p)
	}
	return top
}