/* 229. 求众数 II
给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。

说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。

示例 1:

输入: [3,2,3]
输出: [3]
示例 2:

输入: [1,1,1,3,3,2,2,2]
输出: [1,2] */

func majorityElement(nums []int) []int {
	c := make(map[int]int)
	for _, v := range nums {
		if value, ok := c[v]; ok {
			c[v] = value + 1
		} else {
			c[v] = 1
		}
	}
	r := []int{}
	for k, v := range c {
		if v > len(nums)/3 {
			r = append(r, k)
		}
	}
	return r
}