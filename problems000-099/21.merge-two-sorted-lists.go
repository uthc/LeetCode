/* 合并两个有序链表
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l := new(ListNode)
	start := l
	// if l1==nil{
	//     return l2
	// }
	// if l2 ==nil{
	//     return l1
	// }
	for l1 != nil && l2 != nil {
		l.Next = new(ListNode)
		l = l.Next
		if l1.Val <= l2.Val {
			l.Val = l1.Val
			l1 = l1.Next
		} else {
			l.Val = l2.Val
			l2 = l2.Next
		}
		// if l1==nil&&l2==nil{
		//     return start
		// }

	}
	if l1 != nil {
		l.Next = l1
	}
	if l2 != nil {
		l.Next = l2
	}
	return start.Next
}