/* 合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6 */

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	result := new(ListNode)
	re := result
	for min := 0; ; {
		for min = 0; min < len(lists) && lists[min] == nil; min++ {
		}
		if min == len(lists) {
			break
		}
		for i := 0; i < len(lists); i++ {
			if lists[i] == nil {
				continue
			}
			if lists[min].Val >= lists[i].Val {
				min = i
			}
		}
		re.Next = lists[min]
		lists[min] = lists[min].Next
		re = re.Next
	}
	return result.Next
}