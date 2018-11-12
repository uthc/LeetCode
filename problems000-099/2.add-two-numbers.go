/* 两数相加
给定两个非空链表来表示两个非负整数。位数按照逆序方式存储，它们的每个节点只存储单个数字。将两数相加返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var result *ListNode = new(ListNode)
	xl := result
	var carry int = 0
	for l1 != nil && l2 != nil {
		xl.Next = new(ListNode)
		xl = xl.Next
		xl.Val = (l1.Val + l2.Val + carry) % 10
		carry = (l1.Val + l2.Val + carry) / 10
		l1 = l1.Next
		l2 = l2.Next
	}
	for l1 != nil {
		xl.Next = new(ListNode)
		xl = xl.Next
		xl.Val = (l1.Val + carry) % 10
		carry = (l1.Val + carry) / 10
		l1 = l1.Next
	}
	for l2 != nil {
		xl.Next = new(ListNode)
		xl = xl.Next
		xl.Val = (l2.Val + carry) % 10
		carry = (l2.Val + carry) / 10
		l2 = l2.Next
	}
	if carry != 0 {
		xl.Next = new(ListNode)
		xl = xl.Next
		xl.Val = carry
		carry = 0
	}
	result = result.Next
	return result
}