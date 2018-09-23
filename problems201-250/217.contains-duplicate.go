/* 217. 存在重复元素
题目描述提示帮助提交记录社区讨论阅读解答
给定一个整数数组，判断是否存在重复元素。

如果任何值在数组中出现至少两次，函数返回 true。如果数组中每个元素都不相同，则返回 false。

示例 1:

输入: [1,2,3,1]
输出: true
示例 2:

输入: [1,2,3,4]
输出: false
示例 3:

输入: [1,1,1,3,3,4,3,2,4,2]
输出: true
您是否在真实的面试环节中遇到过这道题目呢？
题目难度：简单
通过次数：22.5K
提交次数：54.8K
贡献者：LeetCode
相关话题

相似题目
存在重复元素 II存在重复元素 III
*/
func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, v := range nums {
		exist, _ := m[v]
		if exist {
			return true
		}
		m[v] = true
	}
	return false
}