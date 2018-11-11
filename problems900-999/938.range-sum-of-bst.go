/* 二叉搜索树的范围和
给定二叉搜索树的根结点 root，返回 L 和 R（含）之间的所有结点的值的和。

二叉搜索树保证具有唯一的值。



示例 1：

输入：root = [10,5,15,3,7,null,18], L = 7, R = 15
输出：32
示例 2：

输入：root = [10,5,15,3,7,13,18,1,null,6], L = 6, R = 10
输出：23


提示：

树中的结点数量最多为 10000 个。
最终的答案保证小于 2^31。 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func rangeSumBST(root *TreeNode, L int, R int) int {
	return getV(root, L, R)
}

func getV(root *TreeNode, L int, R int) int {
	if root == nil {
		return 0
	}
	if root.Val > R {
		return getV(root.Left, L, R)
	}
	if root.Val < L {
		return getV(root.Right, L, R)
	}
	return root.Val + getV(root.Left, L, R) + getV(root.Right, L, R)
}

func findV(root *TreeNode, target int) *TreeNode {
	for root != nil && root.Val != target {
		if root.Val > target {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}