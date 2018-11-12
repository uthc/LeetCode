/* 110. 平衡二叉树
给定一个二叉树，判断它是否是高度平衡的二叉树。

本题中，一棵高度平衡二叉树定义为：

一个二叉树每个节点 的左右两个子树的高度差的绝对值不超过1。

示例 1:

给定二叉树 [3,9,20,null,null,15,7]

    3
   / \
  9  20
    /  \
   15   7
返回 true 。

示例 2:

给定二叉树 [1,2,2,3,3,null,null,4,4]

       1
      / \
     2   2
    / \
   3   3
  / \
 4   4 */

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	sub := childLength(root.Left) - childLength(root.Right)
	if sub > 1 || sub < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func childLength(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l, r := childLength(root.Left), childLength(root.Right)
	if l >= r {
		return l + 1
	}
	return r + 1
}
